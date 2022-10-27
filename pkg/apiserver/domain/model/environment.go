package model

import "kubegems.io/kubegems/pkg/apiserver/options"

type Environment struct {
	ID             uint
	Name           string
	Namespace      string
	Remark         string
	MetaType       string
	DeletePolicy   string
	Cluster        *Cluster
	Project        *Project
	Creator        *User
	ProjectID      uint
	ClusterID      uint
	CreatorID      uint
	Applications   []*Application
	Users          []*User
	ResourceQuota  []byte
	LimitRange     []byte
	VirtualSpaceID *uint
	VirtualSpace   *VirtualSpace
}

func (env *Environment) GetID() uint {
	return env.ID
}

func (env *Environment) GetKind() string {
	return "environment"
}

func (env *Environment) GetName() string {
	return env.Name
}

func (rel *EnvironmentUserRel) SetRole(role string) {
	rel.Role = role
}
func (rel *EnvironmentUserRel) SetUserID(uid uint) {
	rel.UserID = uid
}
func (rel *EnvironmentUserRel) SetInstanceID(envID uint) {
	rel.EnvironmentID = envID
}

func (rel *EnvironmentUserRel) Condition() []options.Option {
	opts := []options.Option{
		options.Equal("user_id", rel.UserID),
		options.Equal("application_id", rel.EnvironmentID),
	}
	return opts
}

func (rel *EnvironmentUserRel) GetRole() string {
	return rel.Role
}

func (rel *EnvironmentUserRel) GetInstanceId() uint {
	return rel.EnvironmentID
}

func (rel *EnvironmentUserRel) GetUserID() uint {
	return rel.UserID
}

type EnvironmentUserRel struct {
	ID            uint
	User          *User
	Environment   *Environment
	UserID        uint
	EnvironmentID uint
	Role          string
}
