package service

import (
	"context"
	"errors"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type TenantManager interface {
	CreateTenant(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error)
	GetTenant(ctx context.Context, name string) (*model.Tenant, error)
	ListTenant(ctx context.Context, opts ...options.Option) ([]*model.Tenant, error)
	DeleteTenant(ctx context.Context, opts ...options.Option) error
	ModifyTenant(ctx context.Context, name string, tenant *model.Tenant) error
}

type tenantManager struct {
	repo repository.GenericRepo[*model.Tenant]
}

func (mt *tenantManager) CreateTenant(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error) {
	exist, err := mt.repo.Get(options.Equal("name", tenant.Name))
	if err == nil {
		return exist, errors.New("exsit")
	}
	// TODO
	if err != nil {
		println(err.Error())
	}
	err = mt.repo.Create(tenant)
	return tenant, err
}

func (mt *tenantManager) GetTenant(ctx context.Context, name string) (*model.Tenant, error) {
	return mt.repo.Get(options.Equal("name", name))
}

func (mt *tenantManager) ListTenant(ctx context.Context, opts ...options.Option) ([]*model.Tenant, error) {
	return mt.repo.List()
}

func (mt *tenantManager) DeleteTenant(ctx context.Context, opts ...options.Option) error {
	return mt.repo.Delete(opts...)
}

func (mt *tenantManager) ModifyTenant(ctx context.Context, name string, tenant *model.Tenant) error {
	exist, err := mt.repo.Get(options.Equal("name", name))
	if exist == nil {
		return errors.New("can't modify not exist tenant")
	}
	// TODO
	if err != nil {
		println(err.Error())
	}
	exist.Remark = tenant.Remark
	exist.Enabled = tenant.Enabled
	return mt.repo.Save(exist)
}

func NewTenantManager(tenantRepo repository.GenericRepo[*model.Tenant]) *tenantManager {
	return &tenantManager{
		repo: tenantRepo,
	}
}
