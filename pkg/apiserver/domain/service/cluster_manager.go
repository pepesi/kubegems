package service

import (
	"context"
	"errors"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type ClusterManager interface {
	CreateCluster(ctx context.Context, tenant *model.Cluster) (*model.Cluster, error)
	GetCluster(ctx context.Context, name string) (*model.Cluster, error)
	ListCluster(ctx context.Context, opts ...options.Option) ([]*model.Cluster, error)
	DeleteCluster(ctx context.Context, opts ...options.Option) error
	ModifyCluster(ctx context.Context, name string, tenant *model.Cluster) error
}

type clusterManager struct {
	repo repository.GenericRepo[*model.Cluster]
}

func (mt *clusterManager) CreateCluster(ctx context.Context, cluster *model.Cluster) (*model.Cluster, error) {
	exist, err := mt.repo.Get(options.Equal("name", cluster.Name))
	if err == nil {
		return exist, errors.New("exsit")
	}
	err = mt.repo.Create(cluster)
	return cluster, err
}

func (mt *clusterManager) GetCluster(ctx context.Context, name string) (*model.Cluster, error) {
	return mt.repo.Get(options.Equal("name", name))
}

func (mt *clusterManager) ListCluster(ctx context.Context, opts ...options.Option) ([]*model.Cluster, error) {
	return mt.repo.List()
}

func (mt *clusterManager) DeleteCluster(ctx context.Context, opts ...options.Option) error {
	return mt.repo.Delete(opts...)
}

func (mt *clusterManager) ModifyCluster(ctx context.Context, name string, tenant *model.Cluster) error {
	exist, err := mt.repo.Get(options.Equal("name", name))
	if exist == nil {
		return errors.New("can't modify not exist tenant")
	}
	// TODO
	if err != nil {
		println(err.Error())
	}
	exist.Name = tenant.Name
	return mt.repo.Save(exist)
}

func NewClusterManager(clusterRepo repository.GenericRepo[*model.Cluster]) *clusterManager {
	return &clusterManager{
		repo: clusterRepo,
	}
}
