package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	domain "kubegems.io/kubegems/pkg/apiserver/domain/service"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type TenantProjectService struct {
	tenantRepo  repository.GenericRepo[*model.Tenant]
	projectRepo repository.GenericRepo[*model.Project]
}

// 项目聚合服务
// 项目管理 crud
// 项目成员管理
// 项目配额管理

func (s *TenantProjectService) AddProject(ctx context.Context, tenantName string, req *v1.Project) (*model.Project, error) {
	tenant, err := s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return nil, err
	}
	mgr := domain.TenantProjectManagerFor(tenant, s.projectRepo)
	proj := &model.Project{
		TenantID: tenant.GetID(),
		Name:     req.Name,
		Remark:   req.Intro,
	}
	err = mgr.CreateProject(proj)
	return proj, err
}

func (s *TenantProjectService) ModifyProject(tenantName string, project *model.Project) error {
	tenant, err := s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return err
	}
	exist, err := s.projectRepo.Get()
	if err != nil {
		return err
	}
	// more fields
	exist.Remark = project.Remark
	mgr := domain.TenantProjectManagerFor(tenant, s.projectRepo)
	return mgr.ModifyProject(exist)
}
