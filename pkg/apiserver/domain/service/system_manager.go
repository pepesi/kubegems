package service

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type UserManager interface {
}

type UserTenantManager interface {
	ListJoinedTenants()
	ListJoinedProjects()
	ListJoinedApplication()
	IsAdmin() bool
}

type AnnouncementManager interface {
	CreateAnnouncement(ctx context.Context, id uint) (*model.Announcement, error)
	GetAnnoucement(ctx context.Context, id uint) (*model.Announcement, error)
	ListAnnouncement(ctx context.Context, opts ...options.Option) ([]*model.Announcement, error)
	DeleteAnnouncement(ctx context.Context, id uint) error
	DisableAnnouncement(ctx context.Context, id uint) error
	EnableAnnouncement(ctx context.Context, id uint) error
}

type announcementManager struct {
	repo repository.GenericRepo[*model.Announcement]
}

func NewAnnouncementManager(repo repository.GenericRepo[*model.Announcement]) AnnouncementManager {
	return &announcementManager{
		repo: repo,
	}
}

func (a *announcementManager) CreateAnnouncement(ctx context.Context, id uint) (*model.Announcement, error) {
	return a.repo.Get(options.Equal("id", id))
}

func (a *announcementManager) GetAnnoucement(ctx context.Context, id uint) (*model.Announcement, error) {
	return a.repo.Get(options.Equal("id", id))
}

func (a *announcementManager) ListAnnouncement(ctx context.Context, opts ...options.Option) ([]*model.Announcement, error) {
	return a.repo.List()
}

func (a *announcementManager) DeleteAnnouncement(ctx context.Context, id uint) error {
	return a.repo.Delete(options.Equal("id", id))
}

func (a *announcementManager) DisableAnnouncement(ctx context.Context, id uint) error {
	anno, err := a.repo.Get()
	if err != nil {
		return err
	}
	if !anno.Enabled {
		return nil
	}
	anno.Enabled = false
	return a.repo.Update(anno)
}

func (a *announcementManager) EnableAnnouncement(ctx context.Context, id uint) error {
	anno, err := a.repo.Get()
	if err != nil {
		return err
	}
	if anno.Enabled {
		return nil
	}
	anno.Enabled = true
	return a.repo.Update(anno)
}

var _ AnnouncementManager = &announcementManager{}
