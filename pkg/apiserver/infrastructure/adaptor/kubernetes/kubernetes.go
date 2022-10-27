package kubernetes

import "sigs.k8s.io/controller-runtime/pkg/client"

type KubeResourceSynchronizer interface {
	Get(client.ObjectKey, client.Object) error
	Create(client.Object) error
	Update(client.Object) error
	Delete(client.ObjectKey) error
}
