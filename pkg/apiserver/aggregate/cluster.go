package agg_services

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/service"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

type ClusterService struct {
	clusterMgr service.ClusterManager
}

func NewClusterService(mgr service.ClusterManager) *ClusterService {
	return &ClusterService{
		clusterMgr: mgr,
	}
}

// 集群服务

// 集群的管理
// 集群插件管理

// AddClusterViaKubeConfig 以kubeconfig方式添加集群
func (s *ClusterService) AddClusterViaKubeConfig(name, cfg string) error {
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
func (s *ClusterService) DeleteCluster(ctx context.Context, name string) error {
	cluster, err := s.clusterMgr.GetCluster(ctx, name)
	if err != nil {
		return err
	}
	return s.clusterMgr.DeleteCluster(ctx, options.Equal("id", cluster.ID))
}

// AddClusterViaRegist 以注册方式添加集群
func (s *ClusterService) AddClusterViaRegist(ctx context.Context, name string) {
}

// RegistCluster 处理集群注册
func (s *ClusterService) RegistCluster(name, cfg string) {}

// EnablePlugin 开启插件
func (s *ClusterService) EnablePlugin(name, plugin string) {}

// DisablePlugin 关闭插件
func (s *ClusterService) DisablePlugin(name, plugin string) {}
