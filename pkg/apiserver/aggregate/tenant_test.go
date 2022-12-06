package agg_services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTenantService_AddMember(t *testing.T) {
	s := NewTenantService(tenantMgr, tenantRepo, tenantRelRepo, userRepo, clusterRepo, quotaRepo)
	Convey("Give tenant name, user name and role", t, func() {
		Convey("When tenant and user both exists", func() {
			tenantName := "kubegems"
			userName := "pepesi"
			role := "admin"
			Convey("Then the result should be nil", func() {
				So(s.AddMember(tenantName, userName, role), ShouldBeNil)
			})
		})
		Convey("When tenant not exist but user exists", func() {
			tenantName := "fake_kubegems"
			userName := "pepesi"
			role := "admin"
			Convey("Then the result should be error", func() {
				So(s.AddMember(tenantName, userName, role), ShouldBeError)
			})
		})
		Convey("When tenant exist but user not exists", func() {
			tenantName := "kubegems"
			userName := "fake_pepesi"
			role := "admin"
			Convey("Then the result should be error", func() {
				So(s.AddMember(tenantName, userName, role), ShouldBeError)
			})
		})
		Convey("When both user and tenant not exists", func() {
			tenantName := "fake_kubegems"
			userName := "fake_pepesi"
			role := "admin"
			Convey("Then the result should be error", func() {
				So(s.AddMember(tenantName, userName, role), ShouldBeError)
			})
		})
	})
}

func TestTenantService_RemoveMember(t *testing.T) {
	s := NewTenantService(tenantMgr, tenantRepo, tenantRelRepo, userRepo, clusterRepo, quotaRepo)
	Convey("Give tenant name and user name", t, func() {
		Convey("When tenant not exist and user exist", func() {
			tenantName := "fake_kubegems"
			userName := "pepesi"
			Convey("Then the result should be error", func() {
				So(s.RemoveMember(tenantName, userName), ShouldBeError)
			})
		})
		Convey("When tenant exist but user not exist", func() {
			tenantName := "kubegems"
			userName := "fake_pepesi"
			Convey("Then the result should be error", func() {
				So(s.RemoveMember(tenantName, userName), ShouldBeError)
			})
		})
		Convey("When both tenant and user exists, but user not in tenant members", func() {
			tenantName := "kubegems"
			userName := "pepesi"
			Convey("Then the result should be nil", func() {
				So(s.RemoveMember(tenantName, userName), ShouldBeNil)
			})
		})
		Convey("When both tenant and user exists, user is tenant member", func() {
			tenantName := "kubegems"
			userName := "cola"
			Convey("Then the result should be nil", func() {
				So(s.RemoveMember(tenantName, userName), ShouldBeNil)
			})
		})
	})
}

func TestTenantService_UpdateMemberRole(t *testing.T) {
	s := NewTenantService(tenantMgr, tenantRepo, tenantRelRepo, userRepo, clusterRepo, quotaRepo)
	Convey("Give tenant name and user name", t, func() {
		Convey("When tenant not exist and user exist", func() {
			tenantName := "fake_kubegems"
			userName := "pepesi"
			Convey("Then the result should be error", func() {
				So(s.UpdateMemberRole(tenantName, userName, "admin"), ShouldBeError)
			})
		})
		Convey("When tenant exist but user not exist", func() {
			tenantName := "kubegems"
			userName := "fake_pepesi"
			Convey("Then the result should be error", func() {
				So(s.UpdateMemberRole(tenantName, userName, ""), ShouldBeError)
			})
		})
		Convey("When user is not tenant member", func() {
			tenantName := "wahaha"
			userName := "pepesi"
			role := "admin"
			Convey("Then the result should be error", func() {
				So(s.UpdateMemberRole(tenantName, userName, role), ShouldBeError)
			})
		})
		Convey("When user is tenant member", func() {
			tenantName := "wahaha"
			userName := "cola"
			role := "admin"
			Convey("Then the result should be nil", func() {
				So(s.UpdateMemberRole(tenantName, userName, role), ShouldBeNil)
			})
		})
	})
}

func TestTenantService_ListMembers(t *testing.T) {
	s := NewTenantService(tenantMgr, tenantRepo, tenantRelRepo, userRepo, clusterRepo, quotaRepo)
	Convey("Give tenant name", t, func() {
		Convey("When tenant exists, list tenant members 2", func() {
			tenantName := "github"
			Convey("Then the result should be 2", func() {
				list, err := s.ListMembers(tenantName)
				So(list, ShouldHaveLength, 2)
				So(err, ShouldBeNil)
			})
		})
		Convey("When tenant exists, list tenant members 1", func() {
			tenantName := "kubegems"
			Convey("Then the result should be 1", func() {
				list, err := s.ListMembers(tenantName)
				So(list, ShouldHaveLength, 1)
				So(err, ShouldBeNil)
			})
		})
		Convey("When tenant not exists, list tenant members", func() {
			tenantName := "fake_tenant"
			Convey("Then the result should be error", func() {
				list, err := s.ListMembers(tenantName)
				So(list, ShouldBeNil)
				So(err, ShouldBeError)
			})
		})
	})
}
