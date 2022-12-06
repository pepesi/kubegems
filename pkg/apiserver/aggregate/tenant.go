package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	domain "kubegems.io/kubegems/pkg/apiserver/domain/service"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
	apis "kubegems.io/kubegems/pkg/apiserver/interfaces/dto/v1"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

// TenantService 1. 租户管理 2. 租户成员 3. 租户集群资源管理
type TenantService interface {
	CreateTenant(ctx context.Context, req *v1.CreateTenantRequest) (*model.Tenant, error)
	GetTenant(ctx context.Context, tenant string) (*apis.RetrieveTenantResp, error)
	ListTenants(ctx context.Context, opts ...options.Option) (*apis.ListTenantResp, error)
	DeleteTenant(ctx context.Context, name string) error
	UpdateTenant(ctx context.Context, name string, req *apis.CreateUpdateTenantReq) error
	AddMember(tenantName, userName, role string) error
	RemoveMember(tenantName, userName string) error
	UpdateMemberRole(tenantName, userName, role string) error
	ListMembers(tenantName string, opts ...options.Option) ([]*model.User, error)
	getTenantUser(tenantName, userName string) (tenant *model.Tenant, user *model.User, err error)
	CreateTenantClusterQuota(tenantName, clusterName string, req *apis.CreateTenantClusterQuotaReq) (*apis.CreateTenantClusterQuotaResp, error)
	ModifyTenantClusterQuota(tenantName, clusterName string, quota model.Quota) error
	getTenantCluster(tenantName, clusterName string) (tenant *model.Tenant, cluster *model.Cluster, err error)
}

type tenantServiceImpl struct {
	tenantManager domain.TenantManager
	tenantRepo    repository.GenericRepo[*model.Tenant]
	tenantRelRepo repository.GenericRepo[*model.TenantUserRel]
	userRepo      repository.GenericRepo[*model.User]
	clusterRepo   repository.GenericRepo[*model.Cluster]
	quotaRepo     repository.GenericRepo[*model.Quota]
}

func NewTenantService(
	mgr domain.TenantManager,
	tenantRepo repository.GenericRepo[*model.Tenant],
	tenantRelRepo repository.GenericRepo[*model.TenantUserRel],
	userRepo repository.GenericRepo[*model.User],
	clusterRepo repository.GenericRepo[*model.Cluster],
	quotaRepo repository.GenericRepo[*model.Quota],
) TenantService {
	return &tenantServiceImpl{
		tenantManager: mgr,
		tenantRepo:    tenantRepo,
		tenantRelRepo: tenantRelRepo,
		userRepo:      userRepo,
		clusterRepo:   clusterRepo,
		quotaRepo:     quotaRepo,
	}
}

func (s *tenantServiceImpl) CreateTenant(ctx context.Context, req *v1.CreateTenantRequest) (*model.Tenant, error) {
	var (
		tenant model.Tenant
	)
	tenant.Name = req.Tenant.Name
	tenant.Remark = req.Tenant.Intro
	return s.tenantManager.CreateTenant(ctx, &tenant)
}

func (s *tenantServiceImpl) GetTenant(ctx context.Context, tenant string) (*apis.RetrieveTenantResp, error) {
	var resp apis.RetrieveTenantResp
	tenantInstance, err := s.tenantManager.GetTenant(ctx, tenant)
	if err != nil {
		return nil, err
	}
	resp.ID = tenantInstance.ID
	resp.Name = tenantInstance.Name
	resp.Remark = tenantInstance.Remark
	resp.CreatedAt = *tenantInstance.CreatedAt
	return &resp, err

}

func (s *tenantServiceImpl) ListTenants(ctx context.Context, opts ...options.Option) (*apis.ListTenantResp, error) {
	var resp apis.ListTenantResp
	tenants, err := s.tenantManager.ListTenant(ctx, opts...)
	if err != nil {
		return nil, err
	}
	for _, tenant := range tenants {
		resp.Tenants = append(resp.Tenants, apis.TenantBase{
			ID:        tenant.ID,
			Name:      tenant.Name,
			Remark:    tenant.Remark,
			Enabled:   tenant.Enabled,
			CreatedAt: *tenant.CreatedAt,
			UpdatedAt: *tenant.UpdatedAt,
		})
	}
	return &resp, nil
}

func (s *tenantServiceImpl) DeleteTenant(ctx context.Context, name string) error {
	return s.tenantManager.DeleteTenant(ctx, options.Equal("name", name))
}

func (s *tenantServiceImpl) UpdateTenant(ctx context.Context, name string, req *apis.CreateUpdateTenantReq) error {
	t := &model.Tenant{
		Name:    req.Name,
		Remark:  req.Remark,
		Enabled: req.Enabled,
	}
	return s.tenantManager.ModifyTenant(ctx, name, t)
}

func (s *tenantServiceImpl) AddMember(tenantName, userName, role string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.AddMember(context.Background(), user, role)
}

func (s *tenantServiceImpl) RemoveMember(tenantName, userName string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.RemoveMember(context.Background(), user)
}

func (s *tenantServiceImpl) UpdateMemberRole(tenantName, userName, role string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.ModifyMemberRole(context.Background(), user, role)
}

func (s *tenantServiceImpl) ListMembers(tenantName string, opts ...options.Option) ([]*model.User, error) {
	tenant, err := s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return nil, err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.ListMember(context.Background(), opts...)
}

func (s *tenantServiceImpl) getTenantUser(tenantName, userName string) (tenant *model.Tenant, user *model.User, err error) {
	tenant, err = s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return
	}
	user, err = s.userRepo.Get(options.Equal("username", userName))
	return
}

func (s *tenantServiceImpl) CreateTenantClusterQuota(tenantName, clusterName string, req *apis.CreateTenantClusterQuotaReq) (*apis.CreateTenantClusterQuotaResp, error) {
	tenant, cluster, err := s.getTenantCluster(tenantName, clusterName)
	if err != nil {
		return nil, err
	}
	mgr := domain.QuotaManagerFor(tenant, s.quotaRepo)
	var (
		quota model.Quota
		resp  apis.CreateTenantClusterQuotaResp
	)
	quota.Datas = req.Datas
	quota.RelName = clusterName
	quota.RelID = cluster.ID
	quota.RelKind = "cluster"
	resp.Datas = quota.Datas
	return &resp, mgr.CreateQuota(quota)
}

func (s *tenantServiceImpl) ModifyTenantClusterQuota(tenantName, clusterName string, quota model.Quota) error {
	tenant, _, err := s.getTenantCluster(tenantName, clusterName)
	if err != nil {
		return err
	}
	mgr := domain.QuotaManagerFor(tenant, s.quotaRepo)
	return mgr.UpdateQuota(quota)
}

func (s *tenantServiceImpl) getTenantCluster(tenantName, clusterName string) (tenant *model.Tenant, cluster *model.Cluster, err error) {
	tenant, err = s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return
	}
	cluster, err = s.clusterRepo.Get(options.Equal("name", clusterName))
	return
}
