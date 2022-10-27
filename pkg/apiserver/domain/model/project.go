package model

import (
	"time"

	"kubegems.io/kubegems/pkg/apiserver/options"
)

type Project struct {
	ID              uint
	Name            string
	CreatedAt       *time.Time
	ProjectAlias    string
	Remark          string
	ResourceQuota   []byte
	Tenant          *Tenant
	TenantID        uint
	Applications    []*Application
	Environments    []*Environment
	ImageRegistries []*ImageRegistry
	Users           []*User
}

func (proj *Project) GetID() uint {
	return proj.ID
}

func (proj *Project) GetKind() string {
	return "project"
}

func (proj *Project) GetName() string {
	return proj.Name
}

type ProjectUserRel struct {
	ID        uint
	User      *User
	Project   *Project
	Role      string
	UserID    uint
	ProjectID uint
}

func (rel *ProjectUserRel) SetRole(role string) {
	rel.Role = role
}

func (rel *ProjectUserRel) SetUserID(uid uint) {
	rel.UserID = uid
}

func (rel *ProjectUserRel) SetInstanceID(insID uint) {
	rel.ProjectID = insID
}

func (rel *ProjectUserRel) Condition() []options.Option {
	opts := []options.Option{
		options.Equal("user_id", rel.UserID),
		options.Equal("project_id", rel.ProjectID),
	}
	return opts
}

func (rel *ProjectUserRel) GetRole() string {
	return rel.Role
}

func (rel *ProjectUserRel) GetInstanceId() uint {
	return rel.ProjectID
}

func (rel *ProjectUserRel) GetUserID() uint {
	return rel.UserID
}
