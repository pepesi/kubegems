package model

import "time"

type Announcement struct {
	ID        uint
	Type      string
	Message   string
	Enabled   bool
	StartAt   time.Time
	EndAt     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (anno *Announcement) IsExpired() bool {
	now := time.Now().Unix()
	return anno.StartAt.Unix() < now && anno.EndAt.Unix() > now
}
