package model

import (
	"time"

	"kubegems.io/kubegems/pkg/apiserver/options"
)

type Application struct {
	ID           uint
	Name         string
	Environments []*Environment
	Project      *Project
	ProjectID    uint
	Remark       string
	Creator      *User
	CreatorID    uint
	CreatedAt    *time.Time
}

func (app *Application) GetID() uint {
	return app.ID
}

type ApplicationUserRel struct {
	ID            uint
	User          *User
	Application   *Application
	UserID        uint
	ApplicationID uint
	Role          string
}

func (rel *ApplicationUserRel) SetRole(role string) {
	rel.Role = role
}

func (rel *ApplicationUserRel) SetUserID(uid uint) {
	rel.UserID = uid
}

func (rel *ApplicationUserRel) SetInstanceID(insID uint) {
	rel.ApplicationID = insID
}

func (rel *ApplicationUserRel) Condition() []options.Option {
	opts := []options.Option{
		options.Equal("user_id", rel.UserID),
		options.Equal("application_id", rel.ApplicationID),
	}
	return opts
}

func (rel *ApplicationUserRel) GetRole() string {
	return rel.Role
}

func (rel *ApplicationUserRel) GetInstanceId() uint {
	return rel.ApplicationID
}

func (rel *ApplicationUserRel) GetUserID() uint {
	return rel.UserID
}

type ApplicationEnvironmentConfig struct {
	EnvironmentID uint
	ApplicationID uint
	Environment   *Environment
	Application   *Application
	Config        []byte
}
