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
)

type LogQueryHistory struct {
	ID         uint         `gorm:"primarykey"`
	Cluster    *Cluster     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	ClusterID  uint         `gorm:"default:0;"`
	LabelJSON  string       `gorm:"type:varchar(1024)"`
	FilterJSON string       `gorm:"type:varchar(1024)"`
	LogQL      string       `gorm:"type:varchar(1024)"`
	CreatedAt  sql.NullTime `gorm:"type:timestamp;default:current_timestamp"`
	Creator    *User        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	CreatorID  uint         `gorm:"default:0"`
}

type LogQuerySnapshot struct {
	ID            uint         `gorm:"primarykey"`
	Cluster       *Cluster     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
	ClusterID     uint         `gorm:"default:0"`
	SnapshotName  string       `gorm:"type:varchar(128);not null;"`
	SourceFile    string       `gorm:"type:varchar(128);not null;"`
	SnapshotCount int          `gorm:"default:0;comment:'line count'"`
	DownloadURL   string       `gorm:"type:varchar(512);not null;"`
	StartTime     sql.NullTime `gorm:"type:timestamp;not null;"`
	EndTime       sql.NullTime `gorm:"type:timestamp;not null;"`
	CreateAt      time.Time    `gorm:"type:timestamp;default:current_timestamp;"`
	Creator       *User        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
	CreatorID     uint         `gorm:"default:0"`
}
