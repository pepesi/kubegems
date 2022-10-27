package agg_services

import (
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	domain "kubegems.io/kubegems/pkg/apiserver/domain/service"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type EnvironmentService struct {
	tenantRepo  repository.GenericRepo[*model.Tenant]
	projectRepo repository.GenericRepo[*model.Project]
	envRepo     repository.GenericRepo[*model.Environment]
	quotaRepo   repository.GenericRepo[*model.Quota]
}

// 环境聚合服务
// 环境资源配额管理
// 环境成员管理
// 环境其他配置管理

func (s *EnvironmentService) getTenantProject(tenantName, projectNmae string) (tenant *model.Tenant, project *model.Project, err error) {
	tenant, err = s.tenantRepo.Get(options.Equal("name", tenantName))
	if err != nil {
		return
	}
	project, err = s.projectRepo.Get(options.Equal("name", projectNmae))
	return
}

func (s *EnvironmentService) CreateEnvironment(tenantName, projectName string, env *model.Environment) error {
	tenant, proj, err := s.getTenantProject(tenantName, projectName)
	if err != nil {
		return err
	}
	mgr := domain.ProjectEnvironmentManagerFor(tenant, proj)
	return mgr.CreateEnvironment(env)
}

func (s *EnvironmentService) ModifyEnvironment(tenantName, projectName string, env *model.Environment) error {
	tenant, proj, err := s.getTenantProject(tenantName, projectName)
	if err != nil {
		return err
	}
	exist, err := s.envRepo.Get(options.Equal("name", env.Name))
	if err != nil {
		return err
	}
	exist.Remark = env.Remark
	mgr := domain.ProjectEnvironmentManagerFor(tenant, proj)
	return mgr.ModifyEnvironment(exist)
}

func (s *EnvironmentService) ModifyEnvironmentQuota(tenantName, projectName, envName string, quota *model.Quota) error {
	// todo belong to tenant, belong to project
	env, err := s.envRepo.Get(
		options.Equal("name", envName),
	)
	if err != nil {
		return err
	}
	quotaMgr := domain.QuotaManagerFor(env, s.quotaRepo)
	return quotaMgr.UpdateQuota(*quota)
}
