package service

import (
	"context"
	"errors"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type MemberManager[T model.HasMemberT] interface {
	AddMember(ctx context.Context, user *model.User, role string) error
	ListMember(ctx context.Context, opts ...options.Option) ([]*model.User, error)
	RemoveMember(ctx context.Context, user *model.User) error
	ModifyMemberRole(ctx context.Context, user *model.User, role string) error
}

type memberManager[T model.HasMemberT, R model.RelationShipT] struct {
	instance T
	relRepo  repository.GenericRepo[R]
	userRepo repository.GenericRepo[*model.User]
}

func (mgr *memberManager[T, R]) AddMember(ctx context.Context, user *model.User, role string) error {
	r := model.RelationInstance[R]()
	r.SetUserID(user.ID)
	r.SetRole(role)
	r.SetInstanceID(mgr.instance.GetID())
	_, err := mgr.relRepo.Get(r.Condition()...)
	if err == nil {
		return errors.New("user is already a tenant member")
	}
	return mgr.relRepo.Create(r)
}

func (mgr *memberManager[T, R]) ListMember(ctx context.Context, opts ...options.Option) ([]*model.User, error) {
	relConditions := mgr.listCondition()
	rels, err := mgr.relRepo.List(relConditions...)
	if err != nil {
		return nil, err
	}
	ids := []uint{}
	for _, rel := range rels {
		ids = append(ids, rel.GetUserID())
	}

	conditions := append(opts, options.Equal("id", ids))
	return mgr.userRepo.List(conditions...)
}

func (mgr *memberManager[T, R]) RemoveMember(ctx context.Context, user *model.User) error {
	r, err := mgr.relRepo.Get(mgr.getCondition(user.ID)...)
	if err != nil {
		return err
	}
	return mgr.relRepo.DeleteInstance(r)
}

func (mgr *memberManager[T, R]) ModifyMemberRole(ctx context.Context, user *model.User, role string) error {
	conds := mgr.getCondition(user.ID)
	r, err := mgr.relRepo.Get(conds...)
	if err != nil {
		return err
	}
	r.SetRole(role)
	return mgr.relRepo.Update(r)
}

func (mgr *memberManager[T, R]) listCondition() []options.Option {
	var t T
	switch any(t).(type) {
	case *model.Tenant:
		return []options.Option{options.Equal("tenant_id", mgr.instance.GetID())}
	case *model.Project:
		return []options.Option{options.Equal("project_id", mgr.instance.GetID())}
	case *model.Application:
		return []options.Option{options.Equal("application_id", mgr.instance.GetID())}
	case *model.Environment:
		return []options.Option{options.Equal("environment_id", mgr.instance.GetID())}
	default:
		return []options.Option{}
	}
}

func (mgr *memberManager[T, R]) getCondition(userid uint) []options.Option {
	var (
		t T
	)
	switch any(t).(type) {
	case *model.Tenant:
		return (&model.TenantUserRel{UserID: userid, TenantID: mgr.instance.GetID()}).Condition()
	case *model.Project:
		return (&model.ProjectUserRel{UserID: userid, ProjectID: mgr.instance.GetID()}).Condition()
	case *model.Application:
		return (&model.ApplicationUserRel{UserID: userid, ApplicationID: mgr.instance.GetID()}).Condition()
	case *model.Environment:
		return (&model.EnvironmentUserRel{UserID: userid, EnvironmentID: mgr.instance.GetID()}).Condition()
	default:
		return []options.Option{}
	}
}

func MemberManagerFor[K model.HasMemberT, R model.RelationShipT](k K, relRepo repository.GenericRepo[R], userRepo repository.GenericRepo[*model.User]) MemberManager[K] {
	return &memberManager[K, R]{
		instance: k,
		relRepo:  relRepo,
		userRepo: userRepo,
	}
}

var _ MemberManager[*model.Tenant] = MemberManagerFor[*model.Tenant, *model.TenantUserRel](nil, nil, nil)

var _ MemberManager[*model.Project] = MemberManagerFor[*model.Project, *model.ProjectUserRel](nil, nil, nil)

var _ MemberManager[*model.Environment] = MemberManagerFor[*model.Environment, *model.EnvironmentUserRel](nil, nil, nil)

var _ MemberManager[*model.Application] = MemberManagerFor[*model.Application, *model.ApplicationUserRel](nil, nil, nil)
