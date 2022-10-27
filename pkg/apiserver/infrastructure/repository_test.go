package infrastructure

import (
	"testing"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure/orm"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

func TestNewGormRepoFor(t *testing.T) {
	db := orm.Init()
	repo := NewGormRepoFor[*model.Tenant](db)
	exists, err := repo.List()
	if err != nil {
		t.Error(err)
	}

	t.Log(">>>", exists)
	var m = model.Tenant{
		ID:     1,
		Name:   "test",
		Remark: "xx",
	}
	err = repo.Create(&m)
	if err != nil {
		t.Error(err)
	}
	var m1 = model.Tenant{
		ID:     2,
		Name:   "ddx",
		Remark: "xx",
	}
	repo.Save(&m1)
	exist, err := repo.Get(options.Equal("id", 1))
	if err != nil {
		t.Error(err)
	}
	t.Log("exist >>", exist)
	exist.Name = "ggg"
	repo.Update(exist)
	exist1, err := repo.Get(options.Equal("id", 1))
	if err != nil {
		panic(err)
	}
	t.Log("updated>>", exist1)
	r, c, e := repo.ListWithCount()
	if e != nil {
		panic(e)
	}
	t.Log("total >>", c)
	t.Log("list >> ", r)
	err = repo.Delete(options.Equal("id", 1))
	if err != nil {
		panic(err)
	}
	err = repo.Delete(options.Equal("id", 2))
	if err != nil {
		panic(err)
	}

}
