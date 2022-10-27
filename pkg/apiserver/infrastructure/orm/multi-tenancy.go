// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
)

type Tenant struct {
	ID                    uint         `gorm:"primarykey;autoIncrement;"`
	Name                  string       `gorm:"type:varchar(50);not null;"`
	Remark                string       `gorm:"type:varchar(1024);default:''"`
	Enabled               sql.NullBool `gorm:"default:true"`
	CreatedAt             sql.NullTime `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt             sql.NullTime `gorm:"type:timestamp"`
	Members               []*User      `gorm:"many2many:tenant_user_rels;"`
	ClusterResourceQuotas []*TenantClusterResourceQuota
	Projects              []*Project
}

type TenantUserRel struct {
	ID       uint    `gorm:"primarykey;autoIncrement;"`
	TenantID uint    `gorm:"uniqueIndex:uniq_idx_tenant_user_rel"`
	UserID   uint    `gorm:"uniqueIndex:uniq_idx_tenant_user_rel"`
	Role     string  `gorm:"type:varchar(30);not null;"`
	Tenant   *Tenant `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	User     *User   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
}

type TenantClusterResourceQuota struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;"`
	TenantID  uint           `gorm:"uniqueIndex:uniq_tenant_cluster"`
	ClusterID uint           `gorm:"uniqueIndex:uniq_tenant_cluster"`
	Tenant    *Tenant        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Cluster   *Cluster       `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Content   datatypes.JSON `gorm:"not null;"`
}

type TenantResourceQuotaApplication struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;"`
	Status    string         `gorm:"type:varchar(30);not null;"`
	Creator   *User          `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET NULL;"`
	CreatedAt sql.NullTime   `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt sql.NullTime   `gorm:"type:timestamp"`
	Content   datatypes.JSON `gorm:"not null;"`
	CreatorID uint
}

type Project struct {
	ID              uint           `gorm:"primarykey;autoIncrement;"`
	Name            string         `gorm:"type:varchar(50);not null;uniqueIndex:uniq_idx_tenant_project_name"`
	CreatedAt       sql.NullTime   `gorm:"type:timestamp;default:current_timestamp;"`
	ProjectAlias    string         `gorm:"type:varchar(50);default:'';"`
	Remark          string         `gorm:"type:varchar(1024);default:'';"`
	ResourceQuota   datatypes.JSON `gorm:"comment:'this field is not used now';"`
	Tenant          *Tenant        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	TenantID        uint           `gorm:"uniqueIndex:uniq_idx_tenant_project_name;"`
	Applications    []*Application
	Environments    []*Environment
	ImageRegistries []*ImageRegistry
	Users           []*User `gorm:"many2many:project_user_rels;"`
}

type ProjectUserRel struct {
	ID        uint     `gorm:"primarykey;autoIncrement;"`
	User      *User    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Project   *Project `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Role      string   `gorm:"type:varchar(30)"`
	UserID    uint     `gorm:"uniqueIndex:uniq_idx_project_user_rel"`
	ProjectID uint     `gorm:"uniqueIndex:uniq_idx_project_user_rel"`
}

type ImageRegistry struct {
	ID            uint         `gorm:"primarykey;autoIncrement;"`
	Name          string       `gorm:"type:varchar(50);uniqueIndex:uniq_idx_project_registry;"`
	Address       string       `gorm:"type:varchar(512)"`
	Username      string       `gorm:"type:varchar(50)"`
	Password      string       `gorm:"type:varchar(512)"`
	UpdateTime    sql.NullTime `gorm:"type:timestamp"`
	Project       *Project     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	ProjectID     uint         `grom:"uniqueIndex:uniq_idx_project_registry;"`
	Creator       *User
	CreatorID     uint
	IsDefault     bool
	EnableExtends bool // 是否启用扩展功能，支持harbor等高级仓库
}

type Environment struct {
	ID             uint           `gorm:"primarykey;autoIncrement;"`
	Name           string         `gorm:"type:varchar(50);not null;uniqueIndex:uniq_idx_project_env;index:environment_uniq,unique"`
	Namespace      string         `gorm:"type:varchar(50)"`
	Remark         string         `gorm:"type:varchar(1024)"`
	MetaType       string         `gorm:"type:varchar(20)"`
	DeletePolicy   string         `gorm:"default:'delNamespace'"`
	Cluster        *Cluster       `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Project        *Project       `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Creator        *User          `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
	ProjectID      uint           `gorm:"uniqueIndex:uniq_idx_project_env"`
	ClusterID      uint           `gorm:"not null;"`
	CreatorID      uint           `gorm:"not null;default:0;"`
	Applications   []*Application `gorm:"many2many:application_environment_configs;"`
	Users          []*User        `gorm:"many2many:environment_user_rels;"`
	ResourceQuota  datatypes.JSON `gorm:"not null"`
	LimitRange     datatypes.JSON `gorm:"not null"`
	VirtualSpaceID *uint          `gorm:"default:0"`
	VirtualSpace   *VirtualSpace  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
}

type EnvironmentUserRel struct {
	ID            uint         `gorm:"primarykey;autoIncrement;"`
	User          *User        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Environment   *Environment `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	UserID        uint         `gorm:"uniqueIndex:uniq_idx_env_user_rel;not null;"`
	EnvironmentID uint         `gorm:"uniqueIndex:uniq_idx_env_user_rel;not null;"`
	Role          string       `gorm:"type:varchar(50);not null;"`
}

type ApplicationEnvironmentConfig struct {
	EnvironmentID uint           `gorm:"not null"`
	ApplicationID uint           `gorm:"not null"`
	Environment   *Environment   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Application   *Application   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Config        datatypes.JSON `gorm:"not null"`
}

type Application struct {
	ID           uint           `gorm:"primarykey;autoIncrement"`
	Name         string         `gorm:"type:varchar(50);uniqueIndex:uniq_idx_project_applicationname;"`
	Environments []*Environment `gorm:"many2many:application_environment_configs;"`
	Project      *Project       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID    uint           `gorm:"uniqueIndex:uniq_idx_project_applicationname;"`
	Remark       string         `gorm:"type:varchar(1024);default:''"`
	Creator      *User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	CreatorID    uint           `gorm:"default:0"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:current_timestamp"`
}
