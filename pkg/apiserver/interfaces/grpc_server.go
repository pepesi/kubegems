package interfaces

import (
	"net"

	"google.golang.org/grpc"

	agg_services "kubegems.io/kubegems/pkg/apiserver/aggregate"
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/domain/service"
	"kubegems.io/kubegems/pkg/apiserver/infrastructure"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
	"kubegems.io/kubegems/pkg/apiserver/interfaces/endpoint"
)

func RunGRPC() {
	infraOpt := infrastructure.NewInfraOption(nil)
	tenantRepo := repository.RepoFor(&model.Tenant{}, infraOpt)
	tenantRelRepo := repository.RepoFor(&model.TenantUserRel{}, infraOpt)
	userRepo := repository.RepoFor(&model.User{}, infraOpt)
	clusterRepo := repository.RepoFor(&model.Cluster{}, infraOpt)
	quotaRepo := repository.RepoFor(&model.Quota{}, infraOpt)

	tenantMgr := service.NewTenantManager(tenantRepo)
	clusterMgr := service.NewClusterManager(clusterRepo)

	tenantService := agg_services.NewTenantService(
		tenantMgr,
		tenantRepo,
		tenantRelRepo,
		userRepo,
		clusterRepo,
		quotaRepo,
	)
	clusterService := agg_services.NewClusterService(
		clusterMgr,
	)

	s := grpc.NewServer()

	v1.RegisterClusterServiceServer(s, endpoint.NewClusterServer(clusterService))
	v1.RegisterTenantServiceServer(s, endpoint.NewTenantServer(tenantService))

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
