package model

import (
	"time"
)

type User struct {
	ID       uint
	Username string
}

type UserGroup struct {
	ID       uint
	Username string
}

type UserGroupRel struct {
	UserID  uint
	GroupID uint
	Role    string
}

type VirtualSpace struct{}

type TenantClusterResourceQuota struct {
	ID        uint
	TenantID  uint
	ClusterID uint
	Content   []byte
}

type TenantResourceQuotaApplication struct {
	ID        uint
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Content   []byte
	CreatorID uint
}

type ImageRegistry struct {
	ID            uint
	Name          string
	Address       string
	Username      string
	Password      string
	UpdateTime    *time.Time
	Project       *Project
	ProjectID     uint
	Creator       *User
	CreatorID     uint
	IsDefault     bool
	EnableExtends bool // 是否启用扩展功能，支持harbor等高级仓库
}
