package orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitDatabase(db *gorm.DB) *gorm.DB {
	if err := db.AutoMigrate(
		// system related
		&SystemRole{},
		&User{},
		&UserToken{},
		&AuditLog{},
		&AuthSource{},
		&Announcement{},
		&AppStoreChartRepo{},
		&UserMessageStatus{},
		&Message{},
		&Cluster{},
		// multi-tenancy
		&Tenant{},
		&TenantUserRel{},
		&TenantClusterResourceQuota{},
		&TenantResourceQuotaApplication{},
		&Project{},
		&ProjectUserRel{},
		&ImageRegistry{},
		&Environment{},
		&EnvironmentUserRel{},
		&ApplicationEnvironmentConfig{},
		&Application{},
		// virtual space
		&VirtualSpace{},
		&VirtualSpaceUserRel{},
		&VirtualDomain{},
		// logging
		&LogQueryHistory{},
		&LogQuerySnapshot{},
		// alerting
		&AlertInfo{},
		&AlertMessage{},
		// monitoring
		&Container{},
		&PromqlTplScope{},
		&PromqlTplResource{},
		&PromqlTplRule{},
		&MonitorDashboard{},
		&MonitorDashboardTpl{},
		&EnvironmentResource{},
	); err != nil {
		panic(err)
	}
	return db
}

type GormModel interface {
	*SystemRole |
		*User |
		*UserToken |
		*AuditLog |
		*Cluster |
		*AuthSource |
		*Announcement |
		*AppStoreChartRepo |
		*UserMessageStatus |
		*Message |
		*Tenant |
		*TenantUserRel |
		*TenantClusterResourceQuota |
		*TenantResourceQuotaApplication |
		*Project |
		*ProjectUserRel |
		*ImageRegistry |
		*Environment |
		*EnvironmentUserRel |
		*ApplicationEnvironmentConfig |
		*Application |
		*VirtualSpace |
		*VirtualSpaceUserRel |
		*VirtualDomain |
		*LogQueryHistory |
		*LogQuerySnapshot |
		*AlertInfo |
		*AlertMessage |
		*Container |
		*PromqlTplScope |
		*PromqlTplResource |
		*PromqlTplRule |
		*MonitorDashboard |
		*MonitorDashboardTpl |
		*EnvironmentResource
}
