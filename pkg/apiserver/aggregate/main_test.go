package agg_services

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/domain/service"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure/orm"
)

var (
	db            *sql.DB
	fixtures      *testfixtures.Loader
	tenantRepo    repository.GenericRepo[*model.Tenant]
	tenantRelRepo repository.GenericRepo[*model.TenantUserRel]
	userRepo      repository.GenericRepo[*model.User]
	clusterRepo   repository.GenericRepo[*model.Cluster]
	quotaRepo     repository.GenericRepo[*model.Quota]
	tenantMgr     service.TenantManager
)

func TestMain(m *testing.M) {
	var err error
	testdbname := "test-database.db"

	gormDB, err := gorm.Open(sqlite.Open(testdbname), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	db, err = gormDB.DB()
	if err != nil {
		panic(err)
	}
	orm.InitDatabase(gormDB)

	infraOpts := infrastructure.NewInfraOption(gormDB)
	tenantRepo = repository.RepoFor(&model.Tenant{}, infraOpts)
	tenantRelRepo = repository.RepoFor(&model.TenantUserRel{}, infraOpts)
	userRepo = repository.RepoFor(&model.User{}, infraOpts)
	clusterRepo = repository.RepoFor(&model.Cluster{}, infraOpts)
	tenantMgr = service.NewTenantManager(tenantRepo)

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory("testdata/fixtures"),
	)
	if err != nil {
		panic(err)
	}
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
	exitcode := m.Run()
	if e := os.Remove(testdbname); e != nil {
		panic(e)
	}
	os.Exit(exitcode)
}
