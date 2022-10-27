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
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID            uint         `gorm:"primarykey;autoIncrement;not null;"`
	Username      string       `gorm:"type:varchar(50);uniqueIndex;not null;"`
	Email         string       `gorm:"type:varchar(50);not null;"`
	Phone         string       `gorm:"type:varchar(255);"`
	Password      string       `gorm:"type:varchar(255);not null;"`
	IsActive      sql.NullBool `gorm:"default:true"`
	CreatedAt     *time.Time   `gorm:"type:timestamp;default:current_timestamp"`
	LastLoginAt   *time.Time   `gorm:"type:timestamp;default:current_timestamp"`
	Source        string       `gorm:"type:varchar(50);not null;"`
	SourceVendor  string       `gorm:"type:varchar(50);not null;"`
	SystemRoleID  uint         `gorm:"not null;default:0"`
	SystemRole    *SystemRole  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	JoinedTenants []*Tenant    `gorm:"many2many:tenant_user_rels;"`
}

type UserToken struct {
	ID        uint          `gorm:"primarykey;autoIncrement;not null;"`
	Name      string        `gorm:"type:varchar(50);not null;"`
	Token     string        `gorm:"type:varchar(256);not null;"`
	GrantType string        `gorm:"type:varchar(50);not null;"`
	Scope     string        `gorm:"type:varchar(50);default:''"`
	ExpireAt  *sql.NullTime `gorm:"type:timestamp"`
	User      *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt *time.Time    `gorm:"type:timestamp;default:current_timestamp;"`
	UserID    uint          `gorm:"not null;"`
}

type SystemRole struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"type:varchar(30);uniqueIndex"`
	Role  string `gorm:"type:varchar(30);uniqueIndex"`
	Users []*User
}

type AuditLog struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	Username  string         `gorm:"type:varchar(50);not null;"`
	Tenant    string         `gorm:"type:varchar(50);not null;"`
	Module    string         `gorm:"type:varchar(512);not null;"`
	Name      string         `gorm:"type:varchar(512);"`
	Action    string         `gorm:"type:varchar(255);not null;"`
	ClientIP  string         `gorm:"type:varchar(255);not null;"`
	Labels    datatypes.JSON `gorm:"not null"`
	RawData   datatypes.JSON `gorm:"not null"`
	Success   bool
}

type Cluster struct {
	ID        uint   `gorm:"primarykey;autoIncrement;"`
	Name      string `gorm:"type:varchar(50);uniqueIndex"`
	APIServer string `gorm:"type:varchar(250);uniqueIndex"`
	Vendor    string `gorm:"type:varchar(50);default:selfhosted"`

	KubeConfig           datatypes.JSON `gorm:"not null"`
	ImageRepo            string         `gorm:"type:varchar(255);default:docker.io/kubegems"`
	DefaultStorageClass  string         `gorm:"type:varchar(128);default:local-path"`
	InstallNamespace     string         `gorm:"type:varchar(64);default:kubegems"`
	Version              string         `gorm:"not null;default:''"`
	AgentAddr            string
	AgentCA              string
	AgentCert            string
	AgentKey             string
	Runtime              string         // docker or containerd
	Primary              bool           // 是否主集群
	OversoldConfig       datatypes.JSON // 集群资源超卖设置
	Environments         []*Environment
	TenantResourceQuotas []*TenantClusterResourceQuota
	DeletedAt            gorm.DeletedAt // soft delete
	ClientCertExpireAt   sql.NullTime   `gorm:"type:timestamp"`
}

type AuthSource struct {
	ID        uint             `gorm:"primaryKey;autoIncrement;"`
	Name      string           `gorm:"type:varchar(50);unique;"`
	Kind      string           `gorm:"type:varchar(30);not null;"`
	Vendor    string           `gorm:"type:varchar(30);not null;"`
	Config    AuthSourceConfig `gorm:"type:json;not null;"`
	TokenType string           `gorm:"type:varchar(50);default:'Bearer';"`
	Enabled   bool             `gorm:"default:true"`
	CreatedAt sql.NullTime     `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt sql.NullTime     `gorm:"type:timestamp"`
}

type AuthSourceConfig struct {
	AuthURL     string   `json:"authURL,omitempty"`
	TokenURL    string   `json:"tokenURL,omitempty"`
	UserInfoURL string   `json:"userInfoURL,omitempty"`
	RedirectURL string   `json:"redirectURL,omitempty"`
	AppID       string   `json:"appID,omitempty"`
	AppSecret   string   `json:"appSecret,omitempty"`
	Scopes      []string `json:"scopes,omitempty"`

	// ldap
	Name         string `json:"name,omitempty"`
	LdapAddr     string `json:"ldapaddr,omitempty"`
	BaseDN       string `json:"basedn,omitempty"`
	EnableTLS    bool   `json:"enableTLS,omitempty"`
	Filter       string `json:"filter,omitempty"`
	BindUsername string `json:"binduser,omitempty"`
	BindPassword string `json:"password,omitempty"`
}

func (cfg *AuthSourceConfig) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := AuthSourceConfig{}
	err := json.Unmarshal(bytes, &result)
	*cfg = result
	return err
}

func (cfg AuthSourceConfig) Value() (driver.Value, error) {
	return json.Marshal(cfg)
}

// Announcement table for site announcements
type Announcement struct {
	ID        uint         `gorm:"primarykey;autoIncrement"`
	Type      string       `gorm:"type:varchar(50);not null;"`
	Message   string       `gorm:"type:varchar(1024);not null;"`
	StartAt   sql.NullTime `gorm:"type:timestamp;not null"`
	EndAt     sql.NullTime `gorm:"type:timestamp;not null"`
	CreatedAt sql.NullTime `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt sql.NullTime `gorm:"type:timestamp;autoUpdateTime"`
}

// AppStoreChartRepo table for site's app store
type AppStoreChartRepo struct {
	ID            uint         `gorm:"primarykey"`
	ChartRepoName string       `gorm:"type:varchar(50);uniqueIndex"`
	URL           string       `gorm:"type:varchar(256)"`
	LastSync      sql.NullTime `gorm:"type:timestamp;autoUpdateTime;"`
	SyncStatus    string       `gorm:"type:varchar(30);not null;"`
	SyncMessage   string       `gorm:"type:varchar(512);not null;"`
}

type UserMessageStatus struct {
	ID             uint
	UserID         uint          `gorm:"default:0;"`
	User           *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MessageID      *uint         `gorm:"default:0;"`
	Message        *Message      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsRead         bool          `gorm:"default:false;"`
	AlertMessageID *uint         `gorm:"default:0;"`
	AlertMessage   *AlertMessage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Message struct {
	ID          uint           `gorm:"primarykey"`
	MessageType string         `gorm:"type:varchar(50);"`
	Title       string         `gorm:"type:varchar(255);"`
	Content     datatypes.JSON `gorm:"not null"`
	CreatedAt   sql.NullTime   `gorm:"type:timestamp;autoUpdateTime;"`
}
