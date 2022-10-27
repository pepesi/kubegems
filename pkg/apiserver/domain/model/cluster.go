package model

import (
	"time"

	"k8s.io/client-go/tools/clientcmd"
	k8s_adaptor "kubegems.io/kubegems/pkg/apiserver/infrastructure/adaptor/kubernetes"
)

type Cluster struct {
	ID        uint
	Name      string
	APIServer string
	Vendor    string

	ImageRepo            string
	DefaultStorageClass  string
	InstallNamespace     string
	Version              string
	AgentAddr            string
	AgentCA              string
	AgentCert            string
	AgentKey             string
	Runtime              string
	Primary              bool
	OversoldConfig       []byte
	KubeConfig           []byte
	Environments         []*Environment
	TenantResourceQuotas []*TenantClusterResourceQuota
	DeletedAt            *time.Time
	ClientCertExpireAt   *time.Time
}

func (cluster *Cluster) CheckConnection() error {
	config, err := clientcmd.NewClientConfigFromBytes(cluster.KubeConfig)
	if err != nil {
		return err
	}
	restConfig, err := config.ClientConfig()
	if err != nil {
		return err
	}
	ck := k8s_adaptor.NewConnectionChecker(restConfig)
	if err := ck.CheckConneciton(); err != nil {
		return err
	}
	return nil
}
