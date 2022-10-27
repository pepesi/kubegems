package model

import (
	"time"

	"kubegems.io/kubegems/pkg/apiserver/options"
)

type Tenant struct {
	ID        uint
	Name      string
	Remark    string
	Enabled   bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (t *Tenant) GetID() uint {
	return t.ID
}

func (t *Tenant) GetKind() string {
	return "tenant"
}

func (t *Tenant) GetName() string {
	return t.Name
}

type TenantUserRel struct {
	ID       uint
	TenantID uint
	UserID   uint
	Role     string
	Tenant   *Tenant
	User     *User
}

func (rel *TenantUserRel) SetRole(role string) {
	rel.Role = role
}

func (rel *TenantUserRel) SetUserID(uid uint) {
	rel.UserID = uid
}

func (rel *TenantUserRel) SetInstanceID(insID uint) {
	rel.TenantID = insID
}

func (rel *TenantUserRel) GetRole() string {
	return rel.Role
}

func (rel *TenantUserRel) GetInstanceId() uint {
	return rel.TenantID
}

func (rel *TenantUserRel) GetUserID() uint {
	return rel.UserID
}

func (rel *TenantUserRel) Condition() []options.Option {
	opts := []options.Option{
		options.Equal("user_id", rel.UserID),
		options.Equal("tenant_id", rel.TenantID),
	}
	return opts
}
