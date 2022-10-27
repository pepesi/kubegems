package infrastructure

import (
	"gorm.io/gorm"
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure/convertor"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

func buildCondition(opts ...options.Option) map[string]interface{} {
	cond := map[string]interface{}{}
	for _, opt := range opts {
		opt.Apply(cond)
	}
	return cond
}

type RepoImpl[M model.Model] struct {
	db *gorm.DB
}

func NewGormRepoFor[M model.Model](db *gorm.DB) *RepoImpl[M] {
	return &RepoImpl[M]{db: db}
}

func (impl *RepoImpl[M]) Exist(opts ...options.Option) bool {
	var count int64
	ins := convertor.NewInfraInstance[M]()
	cond := buildCondition(opts...)
	err := impl.db.Model(ins).Where(cond).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (impl *RepoImpl[M]) Get(opts ...options.Option) (M, error) {
	ins := convertor.NewInfraInstance[M]()
	cond := buildCondition(opts...)
	err := impl.db.First(ins, cond).Error
	r := convertor.Infra2Domain[M](ins)
	return r, err
}

func (impl *RepoImpl[M]) Create(m M) error {
	infraIns := convertor.Domain2GormInfra(m)
	return impl.db.Create(infraIns).Error
}

func (impl *RepoImpl[M]) Save(m M) error {
	infraIns := convertor.Domain2GormInfra(m)
	return impl.db.Save(infraIns).Error
}

func (impl *RepoImpl[M]) List(opts ...options.Option) ([]M, error) {
	infraInsList := convertor.NewInfraInstanceList[M]()
	cond := buildCondition(opts...)
	err := impl.db.Find(infraInsList, cond).Error
	r := convertor.Infra2DomainList[M](infraInsList)
	return r, err
}

func (impl *RepoImpl[M]) ListWithCount(opts ...options.Option) ([]M, int, error) {
	infraInsList := convertor.NewInfraInstanceList[M]()
	cond := buildCondition(opts...)
	var total int64
	impl.db.Model(infraInsList).Where(cond).Count(&total)
	err := impl.db.Find(infraInsList, cond).Error
	r := convertor.Infra2DomainList[M](infraInsList)
	return r, int(total), err
}

func (impl *RepoImpl[M]) Delete(opts ...options.Option) error {
	cond := buildCondition(opts...)
	infraIns := convertor.NewInfraInstance[M]()
	return impl.db.Delete(infraIns, cond).Error
}

func (impl *RepoImpl[M]) DeleteInstance(m M) error {
	infraIns := convertor.Domain2GormInfra(m)
	return impl.db.Delete(infraIns).Error
}

func (impl *RepoImpl[M]) Update(m M) error {
	infraIns := convertor.Domain2GormInfra(m)
	return impl.db.Save(infraIns).Error
}

type InfraOptions struct {
	DB *gorm.DB
}

func NewInfraOption(db *gorm.DB) *InfraOptions {
	return &InfraOptions{
		DB: db,
	}
}
