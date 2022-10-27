package service

import "kubegems.io/kubegems/pkg/apiserver/domain/model"

type ProjectEnvironmentManager interface {
	CreateEnvironment(*model.Environment) error
	ModifyEnvironment(*model.Environment) error
}

type projectEnvironmentManager struct {
	tenant  *model.Tenant
	project *model.Project
}

func (mgr *projectEnvironmentManager) CreateEnvironment(*model.Environment) error {
	return nil
}

func (mgr *projectEnvironmentManager) ModifyEnvironment(*model.Environment) error {
	return nil
}

func ProjectEnvironmentManagerFor(tenant *model.Tenant, project *model.Project) ProjectEnvironmentManager {
	return &projectEnvironmentManager{
		tenant:  tenant,
		project: project,
	}
}
