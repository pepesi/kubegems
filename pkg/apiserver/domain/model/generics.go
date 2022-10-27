package model

import "kubegems.io/kubegems/pkg/apiserver/options"

type Model interface {
	*User | *Cluster | *Quota | *Tenant | *Project | *Environment | *Announcement | *Application |
		*TenantUserRel | *ProjectUserRel | *ApplicationUserRel | *EnvironmentUserRel | *ApplicationEnvironmentConfig
}

type HasMemberT interface {
	GetID() uint
	*Tenant | *Project | *Environment | *Application
}

type RelationShipT interface {
	SetRole(role string)
	SetUserID(uid uint)
	SetInstanceID(insId uint)
	GetRole() string
	GetUserID() uint
	GetInstanceId() uint
	Condition() []options.Option
	*TenantUserRel | *ProjectUserRel | *EnvironmentUserRel | *ApplicationUserRel
}

func RelationInstance[T RelationShipT]() T {
	var (
		t   T
		tmp interface{}
	)
	switch any(t).(type) {
	case *TenantUserRel:
		tmp = &TenantUserRel{}
	case *ProjectUserRel:
		tmp = &ProjectUserRel{}
	case *EnvironmentUserRel:
		tmp = &EnvironmentUserRel{}
	case *ApplicationUserRel:
		tmp = &ApplicationUserRel{}
	}
	return tmp.(T)
}

type HasQuotaT interface {
	GetID() uint
	GetName() string
	GetKind() string
	*Tenant | *Project | *Environment
}
