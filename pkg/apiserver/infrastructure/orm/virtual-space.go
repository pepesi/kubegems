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
)

type VirtualSpace struct {
	ID           uint         `gorm:"primarykey"`
	Name         string       `gorm:"type:varchar(50);uniqueIndex"`
	CreatedAt    sql.NullTime `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt    sql.NullTime `gorm:"type:timestamp"`
	Enabled      bool         `gorm:"default:true"`
	Creator      *User        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
	CreatorID    uint         `gorm:"default:0"`
	Users        []*User      `gorm:"many2many:virtual_space_user_rels;"`
	Environments []*Environment
}

type VirtualSpaceUserRel struct {
	ID             uint          `gorm:"primarykey"`
	VirtualSpaceID uint          `gorm:"uniqueIndex:uniq_vspace_user_rel;"`
	UserID         uint          `gorm:"uniqueIndex:uniq_vspace_user_rel;"`
	VirtualSpace   *VirtualSpace `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	User           *User         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Role           string        `gorm:"type:varchar(30)"`
}

type VirtualDomain struct {
	ID        uint         `gorm:"primarykey;autoIncrement;"`
	Name      string       `gorm:"type:varchar(50);uniqueIndex;"`
	CreatedAt sql.NullTime `gorm:"type:timestamp;default:current_timestamp;"`
	UpdatedAt sql.NullTime `gorm:"type:timestamp;auotUpdateTime;"`
	Enabled   bool         `gorm:"default:true;"`
	Creator   *User        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:SET DEFAULT;"`
	CreatorID uint         `gorm:"default:0;"`
}
