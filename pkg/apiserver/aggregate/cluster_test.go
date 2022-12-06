package agg_services

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"kubegems.io/kubegems/pkg/apiserver/domain/service"
)

func TestClusterService_AddClusterViaKubeConfig(t *testing.T) {
	mgr := service.NewClusterManager(clusterRepo)
	s := NewClusterService(mgr)
	Convey("Give a clustername and kubeconfig string", t, func() {
		Convey("When the cluster connection is normal", func() {
			bytes, _ := os.ReadFile("testdata/kubeconfig/kubecfg1.yaml")
			Convey("Then the result should be nil", func() {
				_, err := s.AddClusterViaKubeConfig("local", string(bytes))
				So(err, ShouldBeNil)
			})
		})
		Convey("When the cluster can't connect", func() {
			bytes, _ := os.ReadFile("testdata/kubeconfig/kubecfg2.yaml")
			Convey("Then the result should be error", func() {
				_, err := s.AddClusterViaKubeConfig("local", string(bytes))
				So(err, ShouldBeError)
			})
		})
	})
}
