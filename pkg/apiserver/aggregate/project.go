package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	domain "kubegems.io/kubegems/pkg/apiserver/domain/service"
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

func (s *TenantProjectService) AddProject(ctx context.Context, tenantName string, project *model.Project) error {
	tenant, err := s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return err
	}
	mgr := domain.TenantProjectManagerFor(tenant, s.projectRepo)
	return mgr.CreateProject(project)
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
