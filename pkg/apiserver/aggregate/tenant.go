package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	domain "kubegems.io/kubegems/pkg/apiserver/domain/service"
	apis "kubegems.io/kubegems/pkg/apiserver/interfaces/dto/v1"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type TenantService struct {
	tenantManager domain.TenantManager
}

// 租户聚合
// 租户管理
// 租户成员
// 租户集群资源管理

func NewTenantService(mgr domain.TenantManager) *TenantService {
	return &TenantService{
		tenantManager: mgr,
	}
}

func (s *TenantService) CreateTenant(ctx context.Context, req *apis.CreateUpdateTenantReq) (*apis.CreateUpdateTenantResp, error) {
	var (
		tenant  model.Tenant
		created apis.CreateUpdateTenantResp
	)
	tenant.Name = req.Name
	tenant.Remark = req.Remark
	_, err := s.tenantManager.CreateTenant(ctx, &tenant)
	created.Name = tenant.Name
	created.Remark = tenant.Remark
	return &created, err
}

func (s *TenantService) GetTenant(ctx context.Context, tenant string) (*apis.RetrieveTenantResp, error) {
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

func (s *TenantService) ListTenants(ctx context.Context, opts ...options.Option) (*apis.ListTenantResp, error) {
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

func (s *TenantService) DeleteTenant(ctx context.Context, name string) error {
	return s.tenantManager.DeleteTenant(ctx, options.Equal("name", name))
}

func (s *TenantService) UpdateTenant(ctx context.Context, name string, req *apis.CreateUpdateTenantReq) error {
	t := &model.Tenant{
		Name:    req.Name,
		Remark:  req.Remark,
		Enabled: req.Enabled,
	}
	return s.tenantManager.ModifyTenant(ctx, name, t)
}

type TenantMemberService struct {
	tenantRepo    repository.GenericRepo[*model.Tenant]
	tenantRelRepo repository.GenericRepo[*model.TenantUserRel]
	userRepo      repository.GenericRepo[*model.User]
}

func NewTenantMemberService(tenantRepo repository.GenericRepo[*model.Tenant], tenantRelRepo repository.GenericRepo[*model.TenantUserRel], userRepo repository.GenericRepo[*model.User]) *TenantMemberService {
	return &TenantMemberService{
		tenantRepo:    tenantRepo,
		tenantRelRepo: tenantRelRepo,
		userRepo:      userRepo,
	}

}

func (s *TenantMemberService) AddMember(tenantName, userName, role string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.AddMember(context.Background(), user, role)
}

func (s *TenantMemberService) RemoveMember(tenantName, userName string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.RemoveMember(context.Background(), user)
}

func (s *TenantMemberService) UpdateMemberRole(tenantName, userName, role string) error {
	tenant, user, err := s.getTenantUser(tenantName, userName)
	if err != nil {
		return err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.ModifyMemberRole(context.Background(), user, role)
}

func (s *TenantMemberService) ListMembers(tenantName string, opts ...options.Option) ([]*model.User, error) {
	tenant, err := s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return nil, err
	}
	mgr := domain.MemberManagerFor(tenant, s.tenantRelRepo, s.userRepo)
	return mgr.ListMember(context.Background(), opts...)
}

func (s *TenantMemberService) getTenantUser(tenantName, userName string) (tenant *model.Tenant, user *model.User, err error) {
	tenant, err = s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return
	}
	user, err = s.userRepo.Get(options.Equal("username", userName))
	return
}

type TenantResourceQuotaService struct {
	tenantRepo  repository.GenericRepo[*model.Tenant]
	clusterRepo repository.GenericRepo[*model.Cluster]
	quotaRepo   repository.GenericRepo[*model.Quota]
}

func NewTenantResourceQuotaService(tenantRepo repository.GenericRepo[*model.Tenant], clusterRepo repository.GenericRepo[*model.Cluster], quotaRepo repository.GenericRepo[*model.Quota]) *TenantResourceQuotaService {
	return &TenantResourceQuotaService{
		tenantRepo:  tenantRepo,
		clusterRepo: clusterRepo,
		quotaRepo:   quotaRepo,
	}
}

func (s *TenantResourceQuotaService) CreateTenantClusterQuota(tenantName, clusterName string, req *apis.CreateTenantClusterQuotaReq) (*apis.CreateTenantClusterQuotaResp, error) {
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

func (s *TenantResourceQuotaService) ModifyTenantClusterQuota(tenantName, clusterName string, quota model.Quota) error {
	tenant, _, err := s.getTenantCluster(tenantName, clusterName)
	if err != nil {
		return err
	}
	mgr := domain.QuotaManagerFor(tenant, s.quotaRepo)
	return mgr.UpdateQuota(quota)
}

func (s *TenantResourceQuotaService) getTenantCluster(tenantName, clusterName string) (tenant *model.Tenant, cluster *model.Cluster, err error) {
	tenant, err = s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return
	}
	cluster, err = s.clusterRepo.Get(options.Equal("name", clusterName))
	return
}
