package repository

import (
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type GenericRepo[T model.Model] interface {
	Exist(opts ...options.Option) bool
	Get(opts ...options.Option) (T, error)
	List(opts ...options.Option) ([]T, error)
	Create(T) error
	Save(T) error
	ListWithCount(opts ...options.Option) ([]T, int, error)
	Delete(opts ...options.Option) error
	DeleteInstance(T) error
	Update(T) error
}

func RepoFor[T model.Model](t T, opt *infrastructure.InfraOptions) GenericRepo[T] {
	return infrastructure.NewGormRepoFor[T](opt.DB)
}
