package endpoint

import (
	"context"

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
	cluster, err := s.clusterSvc.AddClusterViaKubeConfig(req.Name, req.Kubeconfig)
	if err != nil {
		return nil, err
	}
	resp := &v1.CreateClusterResponse{
		Id:      uint32(cluster.ID),
		Name:    cluster.Name,
		Succeed: true,
		Message: "success",
	}
	return resp, nil
}

func (s *clusterImpl) DeleteCluster(ctx context.Context, req *v1.DeleteClusterRequest) (*v1.DeleteClusterResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}
	resp := &v1.DeleteClusterResponse{
		Code:    uint32(v1.StatusCode_STATUS_CODE_UNSPECIFIED),
		Message: "ok",
	}
	if err := s.clusterSvc.DeleteCluster(ctx, req.Name); err != nil {
		resp.Code = uint32(v1.StatusCode_STATUS_CODE_BAD_REQUEST)
		resp.Message = err.Error()
	}
	return resp, nil
}
