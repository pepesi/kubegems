package endpoint

import (
	"context"
	"errors"

	agg_services "kubegems.io/kubegems/pkg/apiserver/aggregate"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

func NewClusterServer(clusterService agg_services.ClusterService) v1.ClusterServiceServer {
	return &clusterImpl{
		clusterSvc: clusterService,
	}
}

type clusterImpl struct {
	v1.UnimplementedClusterServiceServer
	clusterSvc agg_services.ClusterService
}

func (s *clusterImpl) CreateCluster(ctx context.Context, req *v1.CreateClusterRequest) (*v1.CreateClusterResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	return nil, errors.New("not implemented")
}

func (s *clusterImpl) DeleteCluster(ctx context.Context, req *v1.DeleteClusterRequest) (*v1.DeleteClusterResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}
	return nil, errors.New("not implemented")
}
