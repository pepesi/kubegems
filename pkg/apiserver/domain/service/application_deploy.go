package service

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type ManifestStore interface {
	Save(tenant, project, env, app string, manifests map[string][]byte) error
}

type ApplicationManifestsManager interface {
	UpdateManifests([]client.Object) error
}

type applicationManifestsManager struct {
	tenantName      string
	projectName     string
	applicationName string
	store           ManifestStore
}

func (mgr *applicationManifestsManager) UpdateManifests(objs []client.Object) error {
	ret := map[string][]byte{}
	for _, obj := range objs {
		gvk := obj.GetObjectKind().GroupVersionKind()
		group := gvk.Group
		kind := gvk.Kind
		name := obj.GetName()
		uniqName := group + "/" + kind + ":" + name
		data, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		ret[uniqName] = data
	}
	return mgr.store.Save(mgr.tenantName, mgr.projectName, mgr.applicationName, mgr.applicationName, ret)
}

func ApplicationManifestsManagerFor(tenantName, projectName, appName string) ApplicationManifestsManager {
	return &applicationManifestsManager{}
}

type ApplicationDeployManager interface {
	ListApplicationDeployHistory(tenantName, projectName, environment, application string)
	DeployApplication(tenantName, projectName, environment, application string)
	RollbackApplication(tenantName, projectName, environment, application string)
}
