package service

import (
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
)

type TenantProjectManager interface {
	CreateProject(project *model.Project) error
	ModifyProject(project *model.Project) error
}

type tenantProjectManager struct {
	tenant      *model.Tenant
	projectRepo repository.GenericRepo[*model.Project]
}

func (mgr *tenantProjectManager) CreateProject(proj *model.Project) error {
	return nil
}

func (mgr *tenantProjectManager) ModifyProject(proj *model.Project) error {
	return nil
}

func TenantProjectManagerFor(tenant *model.Tenant, projectRepo repository.GenericRepo[*model.Project]) TenantProjectManager {
	return &tenantProjectManager{
		tenant:      tenant,
		projectRepo: projectRepo,
	}
}
