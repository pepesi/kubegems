package interfaces

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"

	gw "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

func RunGrpc() {
	server := &serverImpl{}
	s := grpc.NewServer()
	gw.RegisterKubegemsServiceServer(s, server)
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type serverImpl struct {
	gw.UnimplementedKubegemsServiceServer
}

func (s *serverImpl) CreateTenant(ctx context.Context, req *gw.CreateTenantRequest) (*gw.CreateTenantResponse, error) {
	r := &gw.CreateTenantResponse{
		Succeed: false,
		Message: "xxxx",
	}
	return r, nil
}

func (s *serverImpl) DeleteTenant(ctx context.Context, req *gw.DeleteTenantRequest) (*gw.DeleteTenantResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *serverImpl) CreateCluster(ctx context.Context, req *gw.CreateClusterRequest) (*gw.CreateClusterResponse, error) {
	return nil, errors.New("not implemented")
}
