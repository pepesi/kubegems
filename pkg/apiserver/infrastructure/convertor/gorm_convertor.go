package convertor

import (
	"github.com/jinzhu/copier"
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure/orm"
)

/*
对于orm类型，由于声明的领域模型和数据库orm十分相似，可以直接使用copier完成转换
对付复杂类型，例如 存在 CR 的数据，可能嵌套很深才能获取到某个属性，或者是要通过逻辑计算获取，如果copier不能完成的
那么就需要自己重写 orm2domain_xxx 和 domain2orm_xxx 的方法
*/

func Domain2GormInfra[T model.Model](obj T) interface{} {
	var r interface{}
	switch any(obj).(type) {
	case *model.User:
		r = &orm.User{}
	case *model.Cluster:
		r = &orm.Cluster{}
	case *model.Tenant:
		r = &orm.Tenant{}
	case *model.Project:
		r = &orm.Project{}
	case *model.Application:
		r = &orm.Application{}
	case *model.Environment:
		r = &orm.Environment{}
	case *model.TenantUserRel:
		r = &orm.TenantUserRel{}
	}
	if err := copier.Copy(r, obj); err != nil {
		return nil
	}
	return r
}

func Infra2Domain[T model.Model](obj interface{}) T {
	var converted interface{}
	switch obj.(type) {
	case *orm.User:
		converted = &model.User{}
	case *orm.Cluster:
		converted = &model.Cluster{}
	case *orm.Tenant:
		converted = &model.Tenant{}
	case *orm.Project:
		converted = &model.Project{}
	case *orm.Application:
		converted = &model.Application{}
	case *orm.Environment:
		converted = &model.Environment{}
	case *orm.TenantUserRel:
		converted = &model.TenantUserRel{}
	default:
		// log or panic
	}
	if err := copier.Copy(converted, obj); err != nil {
		return nil
	}
	return converted.(T)
}

func NewInfraInstance[T model.Model]() interface{} {
	var t T
	switch any(t).(type) {
	case *model.User:
		return &orm.User{}
	case *model.Cluster:
		return &orm.Cluster{}
	case *model.Tenant:
		return &orm.Tenant{}
	case *model.Project:
		return &orm.Project{}
	case *model.Application:
		return &orm.Application{}
	case *model.Environment:
		return &orm.Environment{}
	case *model.TenantUserRel:
		return &orm.TenantUserRel{}
	default:
		// log or panic
	}
	return nil
}

func NewInfraInstanceList[T model.Model]() interface{} {
	var t T
	switch any(t).(type) {
	case *model.User:
		return &[]*orm.User{}
	case *model.Cluster:
		return &[]*orm.Cluster{}
	case *model.Tenant:
		return &[]*orm.Tenant{}
	case *model.Project:
		return &[]*orm.Project{}
	case *model.Application:
		return &[]*orm.Application{}
	case *model.Environment:
		return &[]*orm.Environment{}
	case *model.TenantUserRel:
		return &[]*orm.TenantUserRel{}
	default:
		// log or panic
	}
	return nil
}

func Infra2DomainList[T model.Model](obj interface{}) []T {
	var domainList interface{}
	switch ins := obj.(type) {
	case *[]*orm.User:
		domainList = make([]*model.User, len(*ins))
	case *[]*orm.Cluster:
		domainList = make([]*model.Cluster, len(*ins))
	case *[]*orm.Tenant:
		domainList = make([]*model.Tenant, len(*ins))
	case *[]*orm.Project:
		domainList = make([]*model.Project, len(*ins))
	case *[]*orm.Application:
		domainList = make([]*model.Application, len(*ins))
	case *[]*orm.Environment:
		domainList = make([]*model.Environment, len(*ins))
	case *[]*orm.TenantUserRel:
		domainList = make([]*model.TenantUserRel, len(*ins))
	}
	if err := copier.Copy(&domainList, obj); err != nil {
		panic(err)
	}
	return domainList.([]T)
}
