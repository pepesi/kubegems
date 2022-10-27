package service

import (
	"context"
	"errors"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
)

type ApplicationManager interface {
	GetApplication(context.Context, string) (*model.Application, error)
	CreateApplication(ctx context.Context, app *model.Application) error
	ModifyApplication(ctx context.Context, app *model.Application) error
}

type applicationManager struct {
	tenant  *model.Tenant
	project *model.Project
	appRepo repository.GenericRepo[*model.Application]
}

func (appMgr *applicationManager) GetApplication(ctx context.Context, name string) (*model.Application, error) {
	if appMgr.project == nil {
		return nil, errors.New("not found")
	}
	return appMgr.appRepo.Get()
}

func (appMgr *applicationManager) CreateApplication(ctx context.Context, app *model.Application) error {
	if appMgr.project == nil {
		return errors.New("not found")
	}
	return appMgr.appRepo.Save(app)
}

func (appMgr *applicationManager) ListApplication(ctx context.Context) ([]*model.Application, error) {
	return appMgr.appRepo.List()
}

func (appMgr *applicationManager) DeleteApplication(ctx context.Context, name string) error {
	return appMgr.appRepo.Delete()
}

func (appMgr *applicationManager) ModifyApplication(ctx context.Context, app *model.Application) error {
	return appMgr.appRepo.Update(app)
}

func ApplicationManagerFor(tenant *model.Tenant, project *model.Project, appRepo repository.GenericRepo[*model.Application]) *applicationManager {
	return &applicationManager{
		tenant:  tenant,
		project: project,
		appRepo: appRepo,
	}
}

var _ ApplicationManager = ApplicationManagerFor(nil, nil, nil)

type ApplicationEnvironmentManager interface {
	AssociateApplicationToEnvironment(app *model.Application, env *model.Environment) error
	DisAssociateApplicationToEnvironment(app *model.Application, env *model.Environment) error
}

type applicationEnvironmentManager struct {
	appEnvRelRepo repository.GenericRepo[*model.ApplicationEnvironmentConfig]
}

func (mgr *applicationEnvironmentManager) AssociateApplicationToEnvironment(app *model.Application, env *model.Environment) error {
	if mgr.appEnvRelRepo.Exist() {
		return errors.New("exists")
	}
	rel := &model.ApplicationEnvironmentConfig{
		Application: app,
		Environment: env,
	}
	return mgr.appEnvRelRepo.Save(rel)
}

func (mgr *applicationEnvironmentManager) DisAssociateApplicationToEnvironment(app *model.Application, env *model.Environment) error {
	exist, err := mgr.appEnvRelRepo.Get()
	if err != nil {
		return err
	}
	return mgr.appEnvRelRepo.DeleteInstance(exist)
}

func ApplicationEnvironmentManagerFor(appEnvRelRepo repository.GenericRepo[*model.ApplicationEnvironmentConfig]) ApplicationEnvironmentManager {
	return &applicationEnvironmentManager{
		appEnvRelRepo: appEnvRelRepo,
	}
}
