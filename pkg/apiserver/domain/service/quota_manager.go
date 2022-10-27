package service

import (
	"errors"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type QuotaManager[T model.HasMemberT] interface {
	CreateQuota(model.Quota) error
	UpdateQuota(model.Quota) error
}

type quotaManager[T model.HasQuotaT] struct {
	instance  T
	quotaRepo repository.GenericRepo[*model.Quota]
}

func (mgr *quotaManager[T]) CreateQuota(quota model.Quota) error {
	opts := []options.Option{
		options.Equal("kind", quota.RelKind),
		options.Equal("name", quota.RelName),
	}
	_, err := mgr.quotaRepo.Get(opts...)
	if err == nil {
		return errors.New("EXIST")
	}
	return mgr.quotaRepo.Create(&quota)
}

func (mgr *quotaManager[T]) UpdateQuota(quota model.Quota) error {
	opts := []options.Option{
		options.Equal("kind", mgr.instance.GetKind()),
		options.Equal("name", mgr.instance.GetName()),
	}
	existOne, err := mgr.quotaRepo.Get(opts...)
	if err != nil {
		return err
	}
	existOne.Datas = quota.Datas
	return mgr.quotaRepo.Update(existOne)
}

func QuotaManagerFor[HasQuotaT model.HasQuotaT](instance HasQuotaT, quotaRepo repository.GenericRepo[*model.Quota]) QuotaManager[HasQuotaT] {
	return &quotaManager[HasQuotaT]{
		instance:  instance,
		quotaRepo: quotaRepo,
	}
}

var _ QuotaManager[*model.Tenant] = QuotaManagerFor[*model.Tenant](nil, nil)
var _ QuotaManager[*model.Environment] = QuotaManagerFor[*model.Environment](nil, nil)
