package orm

import (
	"database/sql"

	"gorm.io/datatypes"
)

type AlertInfo struct {
	Fingerprint     string `gorm:"type:varchar(50);primaryKey"`
	Name            string `gorm:"type:varchar(50);"`
	Namespace       string `gorm:"type:varchar(50);"`
	ClusterName     string `gorm:"type:varchar(50);"`
	TenantName      string `gorm:"type:varchar(50);index"`
	ProjectName     string `gorm:"type:varchar(50);index"`
	EnvironmentName string `gorm:"type:varchar(50);index"`
	Labels          datatypes.JSON

	SilenceStartsAt  sql.NullTime `gorm:"type:timestamp;"`
	SilenceUpdatedAt sql.NullTime `gorm:"type:timestamp;"`
	SilenceEndsAt    sql.NullTime `gorm:"type:timestamp;"`
	SilenceCreator   string       `gorm:"type:varchar(50);"`
}

type AlertMessage struct {
	ID          uint         `gorm:"primaryKey"`
	Fingerprint string       `gorm:"type:varchar(50);"`
	AlertInfo   *AlertInfo   `gorm:"foreignKey:Fingerprint;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Value       string       `gorm:"not null;"`
	Message     string       `gorm:"not null;"`
	StartsAt    sql.NullTime `gorm:"type:timestamp;index"`
	EndsAt      sql.NullTime `gorm:"type:timestamp;"`
	CreatedAt   sql.NullTime `gorm:"type:timestamp;index"`
	Status      string       `gorm:"type:varchar(20);not null;"`
}
