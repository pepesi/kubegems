package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ConnectionChecker struct {
	config *rest.Config
}

func NewConnectionChecker(config *rest.Config) *ConnectionChecker {
	return &ConnectionChecker{
		config: config,
	}
}

func (c *ConnectionChecker) CheckConneciton() error {
	clientSet, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		return err
	}
	_, err = clientSet.ServerVersion()
	if err != nil {
		return err
	}
	return nil
}
