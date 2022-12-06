package interfaces

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"

	gw "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

func RunGrpc() {
	tenantServer := &tenantImpl{}
	clusterServer := &clusterImpl{}
	s := grpc.NewServer()

	gw.RegisterClusterServiceServer(s, clusterServer)
	gw.RegisterTenantServiceServer(s, tenantServer)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type clusterImpl struct {
	gw.UnimplementedClusterServiceServer
}

func (s *clusterImpl) CreateCluster(ctx context.Context, req *gw.CreateClusterRequest) (*gw.CreateClusterResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *clusterImpl) DeleteCluster(ctx context.Context, req *gw.DeleteClusterRequest) (*gw.DeleteClusterResponse, error) {
	return nil, errors.New("not implemented")
}

type tenantImpl struct {
	gw.UnimplementedTenantServiceServer
}

func (s *tenantImpl) CreateTenant(ctx context.Context, req *gw.CreateTenantRequest) (*gw.CreateTenantResponse, error) {
	r := &gw.CreateTenantResponse{
		Succeed: false,
		Message: "xxxx",
	}
	return r, nil
}

func (s *tenantImpl) DeleteTenant(ctx context.Context, req *gw.DeleteTenantRequest) (*gw.DeleteTenantResponse, error) {
	return nil, errors.New("not implemented")
}
