package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/service"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

// ClusterService  1.集群服务 2.集群管理 3.集群插件管理
type ClusterService interface {
	AddClusterViaKubeConfig(name, cfg string) error
	DeleteCluster(ctx context.Context, name string) error
	AddClusterViaRegist(ctx context.Context, name string)
	RegistCluster(name, cfg string)
	EnablePlugin(name, plugin string)
	DisablePlugin(name, plugin string)
}

func NewClusterService(mgr service.ClusterManager) ClusterService {
	return &clusterServiceImpl{
		clusterMgr: mgr,
	}
}

type clusterServiceImpl struct {
	clusterMgr service.ClusterManager
}

// AddClusterViaKubeConfig 以kubeconfig方式添加集群
func (s *clusterServiceImpl) AddClusterViaKubeConfig(name, cfg string) error {
	cluster := &model.Cluster{
		ID:         1,
		Name:       name,
		KubeConfig: []byte(cfg),
	}
	if err := cluster.CheckConnection(); err != nil {
		return err
	}
	_, err := s.clusterMgr.CreateCluster(context.Background(), cluster)
	return err
}

// DeleteCluster 删除集群
func (s *clusterServiceImpl) DeleteCluster(ctx context.Context, name string) error {
	cluster, err := s.clusterMgr.GetCluster(ctx, name)
	if err != nil {
		return err
	}
	return s.clusterMgr.DeleteCluster(ctx, options.Equal("id", cluster.ID))
}

// AddClusterViaRegist 以注册方式添加集群
func (s *clusterServiceImpl) AddClusterViaRegist(ctx context.Context, name string) {
}

// RegistCluster 处理集群注册
func (s *clusterServiceImpl) RegistCluster(name, cfg string) {}

// EnablePlugin 开启插件
func (s *clusterServiceImpl) EnablePlugin(name, plugin string) {}

// DisablePlugin 关闭插件
func (s *clusterServiceImpl) DisablePlugin(name, plugin string) {}
