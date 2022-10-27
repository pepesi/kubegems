package orm

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"gorm.io/datatypes"
	"kubegems.io/kubegems/pkg/utils/prometheus"
)

type Workload struct {
	ID                uint         `gorm:"primarykey"`
	CreatedAt         sql.NullTime `gorm:"type:timestamp;default:current_timestamp"`
	ClusterName       string       `gorm:"type:varchar(50);not null;"`
	Namespace         string       `gorm:"type:varchar(64);not null;"`
	Type              string       `gorm:"type:varchar(50);not null;"`
	Name              string       `gorm:"type:varchar(128);not null;"`
	CPULimitStdvar    float64      `gorm:"column:cpu_limit_stdvar"`
	MemoryLimitStdvar float64      `gorm:"column:mem_limit_stdvar"`
	Containers        []*Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Container struct {
	ID               uint      `gorm:"primarykey;autoIncrement;"`
	Name             string    `gorm:"type:varchar(128);not null;"`
	PodName          string    `gorm:"type:varchar(128);not null;"`
	CPULimits        float64   `gorm:"column:cpu_limits;"`
	MemoryLimits     int64     `gorm:"column:mem_limits;"`
	CPUUsageCore     float64   `gorm:"column:cpu_useage"`
	CPUPercent       float64   `gorm:"column:cpu_persents"`
	MemoryUsageBytes float64   `gorm:"column:mem_useage"`
	MemoryPercent    float64   `gorm:"column:mem_persents"`
	WorkloadID       uint      `gorm:"default:0"`
	Workload         *Workload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
}

type PromqlTplScope struct {
	ID         uint                 `gorm:"primarykey"`
	Name       string               `gorm:"type:varchar(50);uniqueIndex"`
	ShowName   string               `gorm:"type:varchar(50);not null;"`
	Namespaced bool                 `gorm:"default:true"`
	Resources  []*PromqlTplResource `gorm:"foreignKey:ScopeID"`
	CreatedAt  sql.NullTime         `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  sql.NullTime         `gorm:"type:timestamp;autoUpdateTime"`
}

type PromqlTplResource struct {
	ID        uint             `gorm:"primarykey;autoIncrement"`
	Name      string           `gorm:"type:varchar(50);uniqueIndex"`
	ShowName  string           `gorm:"type:varchar(50)"`
	ScopeID   uint             `gorm:"default:0;"`
	Scope     *PromqlTplScope  `gorm:"foreignKey:ScopeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rules     []*PromqlTplRule `json:"rules,omitempty" gorm:"foreignKey:ResourceID"`
	CreatedAt sql.NullTime     `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt sql.NullTime     `gorm:"type:timestamp;autoUpdateTime"`
}

type PromqlTplRule struct {
	ID          uint               `gorm:"primarykey;autoIncrement"`
	Name        string             `gorm:"type:varchar(50);not null'"`
	ShowName    string             `gorm:"type:varchar(50);not null;"`
	Description string             `gorm:"type:varchar(1024);not null;"`
	Expr        string             `gorm:"type:varchar(1024);not null;"`
	Unit        string             `gorm:"type:varchar(50)" json:"unit"`
	Labels      datatypes.JSON     `gorm:"not null;"`
	ResourceID  *uint              `grom:"default:0;"`
	Resource    *PromqlTplResource `gorm:"foreignKey:ResourceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TenantID    *uint              `grom:"default:0;"`
	Tenant      *Tenant            `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	CreatedAt   sql.NullTime       `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   sql.NullTime       `gorm:"type:timestamp;autoUpdateTime"`
}

type MonitorDashboard struct {
	ID            uint           `gorm:"primarykey"`
	Name          string         `gorm:"type:varchar(50)"`
	Step          string         `gorm:"type:varchar(50)"` // 样本间隔，单位秒
	Refresh       string         `gorm:"type:varchar(50)"` // 刷新间隔，eg. 30s, 1m
	Start         string         `gorm:"type:varchar(50)"` // 开始时间，eg. 2022-04-24 06:00:45.241, now, now-30m
	End           string         `gorm:"type:varchar(50)"` // 结束时间
	CreatedAt     sql.NullTime   `gorm:"type:timestamp;default:current_timestamp"`
	Creator       string         `gorm:"type:varchar(50)" json:"creator"` // 创建者
	Graphs        MonitorGraphs  `gorm:"not null;"`
	Variables     datatypes.JSON `gorm:"not null;"`        // 变量
	Template      string         `gorm:"type:varchar(50)"` // 模板名
	EnvironmentID *uint          `gorm:"default:0"`
	Environment   *Environment   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;" json:"environment"`
}

type MonitorDashboardTpl struct {
	Name        string         `gorm:"type:varchar(50);primaryKey;"`
	Description string         `gorm:"type:varchar(1024);"`
	Step        string         `gorm:"type:varchar(50);"` // 样本间隔，单位秒
	Refresh     string         `gorm:"type:varchar(50);"` // 刷新间隔，eg. 30s, 1m
	Start       string         `gorm:"type:varchar(50);"` // 开始时间，eg. 2022-04-24 06:00:45.241, now, now-30m
	End         string         `gorm:"type:varchar(50);"` // 结束时间
	Graphs      MonitorGraphs  `gorm:"not null;"`
	Variables   datatypes.JSON `gorm:"not null;"`
	CreatedAt   sql.NullTime   `gorm:"type:timestamp;default:current_timestamp;"`
	UpdatedAt   sql.NullTime   `gorm:"type:timestamp;autoUpdateTime;"`
}

type EnvironmentResource struct {
	ID                 uint         `gorm:"primarykey"`
	CreatedAt          sql.NullTime `gorm:"type:timestamp;autoUpdateTime;"`
	ClusterName        string
	TenantName         string
	ProjectName        string
	EnvironmentName    string
	MaxCPUUsageCore    float64
	MaxMemoryUsageByte float64
	MinCPUUsageCore    float64
	MinMemoryUsageByte float64
	AvgCPUUsageCore    float64
	AvgMemoryUsageByte float64
	NetworkReceiveByte float64
	NetworkSendByte    float64
	MaxPVCUsageByte    float64
	MinPVCUsageByte    float64
	AvgPVCUsageByte    float64
}

type MonitorGraphs []MetricGraph

func (g MonitorGraphs) GormDataType() string {
	return "json"
}

func (g *MonitorGraphs) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, g)
	}
	return nil
}

func (g MonitorGraphs) Value() (driver.Value, error) {
	return json.Marshal(g)
}

type MetricGraph struct {
	Name                        string `json:"name"`
	*prometheus.PromqlGenerator `json:"promqlGenerator"`
	Expr                        string `json:"expr"`
	Unit                        string `json:"unit"`
}
