package announcement

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type condition interface {
	Page() uint
	Size() uint
	ExtraCond() map[string]interface{}
}

type PageSizeCond struct {
	PageVal uint `form:"page"`
	SizeVal uint `form:"size"`
}

func (pc *PageSizeCond) Page() uint {
	return pc.PageVal
}

func (pc *PageSizeCond) Size() uint {
	return pc.SizeVal
}

type announcementCondition struct {
	PageSizeCond
	Search  string     `form:"search"`
	StartAt *time.Time `form:"start_at"`
	EndAt   *time.Time `form:"end_at"`
}

func (ac *announcementCondition) ExtraCond() map[string]interface{} {
	extracCondition := map[string]interface{}{}
	if len(ac.Search) > 0 {
		extracCondition["search"] = ac.Search
	}
	if ac.StartAt != nil {
		extracCondition["start_at"] = ac.StartAt
	}
	if ac.EndAt != nil {
		extracCondition["end_at"] = ac.EndAt
	}
	return extracCondition
}

type PrimaryKeyCondition struct {
	PK uint `uri:"id" form:"id"`
}

func (ic *PrimaryKeyCondition) ID() uint {
	return ic.PK
}

type AnnouncementEntity struct {
	ID        uint       `json:"id"`
	Type      string     `json:"type"`
	Message   string     `json:"message"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type AnnouncementService interface {
	Get(pk uint) (*AnnouncementEntity, error)
	Create(entity *AnnouncementEntity) error
	List(condition) ([]*AnnouncementEntity, error)
	Delete(entity *AnnouncementEntity) error
	Modify(entity *AnnouncementEntity) error
}

type AnnouncementAPI struct {
	service AnnouncementService
}

func codeInfoFromError(err error) (int, string) {
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return 0, err.Error()
	}
	if me.Number == 1062 {
		return 400, "It already exists in a database."
	}
	return 400, err.Error()
}

func (h *AnnouncementAPI) RetrieveAnnouncement(ctx *gin.Context) {
	q := struct {
		id   uint `uri:"id"`
		page uint `form:"page"`
		size uint `form:"size"`
	}{}
	ctx.Bind(&q)
	announcement, err := h.service.Get(q.id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, announcement)
}

func (h *AnnouncementAPI) ListAnnouncement(ctx *gin.Context) {
	cond := &announcementCondition{}
	ctx.Bind(cond)
	announcements, err := h.service.List(cond)
	if err != nil {
		ctx.JSON(codeInfoFromError(err))
		return
	}
	ctx.JSON(http.StatusOK, announcements)
}

type AnnouncementServiceImpl struct {
	repo AnnouncementRepo
}

func (annocementImpl *AnnouncementServiceImpl) Get(pk uint) (*AnnouncementEntity, error) {
	return annocementImpl.repo.Get(pk)
}

func (annocementImpl *AnnouncementServiceImpl) Create(entity *AnnouncementEntity) error {
	return annocementImpl.repo.Create(entity)
}

func (annocementImpl *AnnouncementServiceImpl) List(cond condition) ([]*AnnouncementEntity, error) {
	return annocementImpl.repo.List(cond)
}

func (annocementImpl *AnnouncementServiceImpl) Delete(entity *AnnouncementEntity) error {
	// before delete other action
	return annocementImpl.repo.Delete(entity)
}
func (annocementImpl *AnnouncementServiceImpl) Modify(entity *AnnouncementEntity) error {
	return annocementImpl.repo.Modify(entity)
}

type AnnouncementRepo interface {
	Get(pk uint) (*AnnouncementEntity, error)
	Create(entity *AnnouncementEntity) error
	List(condition) ([]*AnnouncementEntity, error)
	Delete(entity *AnnouncementEntity) error
	Modify(entity *AnnouncementEntity) error
}

type Announcement struct {
	ID      uint       `gorm:"primarykey" json:"id"`
	Type    string     `gorm:"type:varchar(50);" json:"type"`
	Message string     `json:"message"`
	StartAt *time.Time `json:"startAt"` // 开始时间，默认现在
	EndAt   *time.Time `json:"endAt"`   // 结束时间，默认一天后

	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type AnnouncementRepoImpl struct {
	db *gorm.DB
}

func (rl *AnnouncementRepoImpl) Get(pk uint) (*AnnouncementEntity, error) {
	var anno Announcement
	err := rl.db.First(&anno, pk).Error
	if err != nil {
		return nil, err
	}
	entigy := AnnouncementEntity{
		ID: anno.ID,
	}
	return &entigy, err
}

func (rl *AnnouncementRepoImpl) List(cond condition) (ret []*AnnouncementEntity, err error) {
	err = rl.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(cond.Size())).Offset(int(cond.Page()-1) * int(cond.Size()))
	}).Find(ret, cond.ExtraCond()).Error
	return
}

func (rl *AnnouncementRepoImpl) Create(entity *AnnouncementEntity) error {
	return nil
}
func (rl *AnnouncementRepoImpl) Modify(entity *AnnouncementEntity) error {
	return nil
}
func (rl *AnnouncementRepoImpl) Delete(entity *AnnouncementEntity) error {
	return nil
}

func Regist() {
	_ = AnnouncementAPI{
		service: &AnnouncementServiceImpl{
			repo: &AnnouncementRepoImpl{},
		},
	}
}
