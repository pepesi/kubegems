// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gemsplugin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/semver/v3"
	"helm.sh/helm/v3/pkg/repo"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pluginscommon "kubegems.io/kubegems/pkg/apis/plugins"
	pluginsv1beta1 "kubegems.io/kubegems/pkg/apis/plugins/v1beta1"
)

type PluginVersion struct {
	Name             string                      `json:"name,omitempty"`
	Namespace        string                      `json:"namespace,omitempty"`
	Enabled          bool                        `json:"enabled,omitempty"`
	InstallNamespace string                      `json:"installNamespace,omitempty"`
	Kind             pluginsv1beta1.BundleKind   `json:"kind,omitempty"`
	Description      string                      `json:"description,omitempty"`
	HelathCheck      string                      `json:"helathCheck,omitempty"`
	MainCategory     string                      `json:"mainCategory,omitempty"`
	Category         string                      `json:"category,omitempty"`
	Repository       string                      `json:"repository,omitempty"`
	Version          string                      `json:"version,omitempty"`
	Healthy          bool                        `json:"healthy,omitempty"`
	Required         bool                        `json:"required,omitempty"`
	Requirements     string                      `json:"requirements,omitempty"` // dependecies requirements
	Message          string                      `json:"message,omitempty"`
	Values           pluginsv1beta1.Values       `json:"values,omitempty"`
	Files            map[string]string           `json:"files,omitempty"`
	ValuesFrom       []pluginsv1beta1.ValuesFrom `json:"valuesFrom,omitempty"`
	Priority         int                         `json:"priority,omitempty"`
}

func (pv PluginVersion) ToPlugin() *pluginsv1beta1.Plugin {
	if pv.Kind == "" {
		pv.Kind = pluginsv1beta1.BundleKindTemplate // prefer use template with plugin
	}
	return &pluginsv1beta1.Plugin{
		ObjectMeta: v1.ObjectMeta{
			Name:      pv.Name,
			Namespace: pv.Namespace,
			Annotations: map[string]string{
				pluginscommon.AnnotationCategory:          pv.MainCategory + "/" + pv.Category,
				pluginscommon.AnnotationPluginDescription: pv.Description,
				pluginscommon.AnnotationHealthCheck:       pv.HelathCheck,
				pluginscommon.AnnotationRequired:          strconv.FormatBool(pv.Required),
			},
		},
		Spec: pluginsv1beta1.PluginSpec{
			Kind:             pv.Kind,
			URL:              pv.Repository,
			InstallNamespace: pv.InstallNamespace,
			Chart:            pv.Name,
			Version:          pv.Version,
			Values:           pv.Values,
			ValuesFrom:       pv.ValuesFrom,
		},
	}
}

func PluginVersionFrom(plugin pluginsv1beta1.Plugin) PluginVersion {
	annotations := plugin.Annotations
	if annotations == nil {
		annotations = map[string]string{}
	}
	maincate, cate := parseCategory(annotations)
	required, _ := strconv.ParseBool(annotations[pluginscommon.AnnotationRequired])
	pv := PluginVersion{
		Name:             plugin.Name,
		Namespace:        plugin.Namespace,
		InstallNamespace: plugin.Spec.InstallNamespace,
		Enabled:          plugin.DeletionTimestamp == nil && !plugin.Spec.Disabled,
		Kind:             plugin.Spec.Kind,
		Description:      annotations[pluginscommon.AnnotationPluginDescription],
		HelathCheck:      annotations[pluginscommon.AnnotationHealthCheck],
		MainCategory:     maincate,
		Category:         cate,
		Repository:       plugin.Spec.URL,
		Version:          plugin.Spec.Version,
		Values:           plugin.Spec.Values,
		ValuesFrom:       plugin.Spec.ValuesFrom,
		Required:         required,
	}
	if plugin.Status.Phase == pluginsv1beta1.PhaseInstalled {
		pv.Healthy = true
	} else {
		pv.Message = plugin.Status.Message // display the message on not installed
	}
	return pv
}

func PluginVersionFromRepoChartVersion(repo string, cv *repo.ChartVersion) PluginVersion {
	annotations := cv.Annotations
	if annotations == nil {
		annotations = map[string]string{}
	}

	valsFroms := []pluginsv1beta1.ValuesFrom{}
	if cv.Name != pluginscommon.KubegemsChartGlobal {
		// always inject the global values reference in other plugin
		valsFroms = append(valsFroms, pluginsv1beta1.ValuesFrom{
			Kind:     pluginsv1beta1.ValuesFromKindConfigmap,
			Name:     pluginscommon.KubeGemsGlobalValuesConfigMapName,
			Prefix:   pluginscommon.KubegemsChartGlobal + ".",
			Optional: true,
		})
	}

	for _, val := range strings.Split(annotations[pluginscommon.AnnotationValuesFrom], ",") {
		if val == "" || val == pluginscommon.KubegemsChartGlobal {
			continue
		}
		namespace, name := "", val
		if splits := strings.Split(val, "/"); len(splits) > 1 {
			namespace, name = splits[0], splits[1]
		}
		valsFroms = append(valsFroms, pluginsv1beta1.ValuesFrom{
			Kind:      pluginsv1beta1.ValuesFromKindConfigmap,
			Name:      fmt.Sprintf("kubegems-%s-values", name),
			Namespace: namespace,
			Prefix:    name + ".",
			Optional:  true,
		})
	}
	maincate, cate := parseCategory(annotations)
	required, _ := strconv.ParseBool(annotations[pluginscommon.AnnotationRequired])
	return PluginVersion{
		Name:             cv.Name,
		Repository:       repo,
		Version:          cv.Version,
		Description:      cv.Description,
		ValuesFrom:       valsFroms,
		InstallNamespace: annotations[pluginscommon.AnnotationInstallNamespace],
		Requirements:     annotations[pluginscommon.AnnotationRequirements],
		HelathCheck:      annotations[pluginscommon.AnnotationHealthCheck],
		Required:         required,
		Kind: func() pluginsv1beta1.BundleKind {
			if kind := annotations[pluginscommon.AnnotationRenderBy]; kind != "" {
				return pluginsv1beta1.BundleKind(kind)
			} else {
				return pluginsv1beta1.BundleKindTemplate
			}
		}(),
		MainCategory: maincate,
		Category:     cate,
	}
}

func parseCategory(annotations map[string]string) (string, string) {
	maincate, cate := "other", "unknow"
	full := annotations[pluginscommon.AnnotationCategory]
	if full == "" {
		return maincate, cate
	}
	categories := strings.Split(full, "/")
	if len(categories) == 1 {
		cate = categories[0]
		if oldmaincate := annotations["plugins.kubegems.io/main-category"]; oldmaincate != "" {
			maincate = oldmaincate
		}
	} else if len(categories) > 1 {
		maincate, cate = categories[0], categories[1]
	}
	return maincate, cate
}

func IsPluginChart(cv *repo.ChartVersion) bool {
	annotations := cv.Annotations
	if annotations == nil {
		return false
	}
	b, _ := strconv.ParseBool(annotations[pluginscommon.AnnotationIsPlugin])
	return b
}

type ErrorList []error

func (list ErrorList) Error() string {
	msg := ""
	for _, item := range list {
		msg += item.Error() + ";"
	}
	return msg
}

func CheckDependecy(requirements string, exist PluginVersion) error {
	reqs := ParseRequirements(requirements)
	if req, ok := reqs[exist.Name]; ok {
		existver, err := semver.NewVersion(exist.Version)
		if err != nil {
			// we cant check version so adopt any.
			return nil
		}
		if !req.Check(existver) {
			return fmt.Errorf("version not matched: %s", exist.Version)
		}
		return nil
	}
	// not required
	return nil
}

func CheckDependecies(requirements string, installs map[string]PluginVersion) error {
	reqs := ParseRequirements(requirements)
	var errs ErrorList
	for name, constraint := range reqs {
		installed, ok := installs[name]
		if !ok {
			errs = append(errs, fmt.Errorf("%s not installed,require: %s", name, constraint))
			continue
		}
		ver, err := semver.NewVersion(installed.Version)
		if err != nil {
			continue
		}
		if !constraint.Check(ver) {
			errs = append(errs, fmt.Errorf("%s not meet,require: %s", name, constraint))
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

type Requirement struct {
	Name       string
	Constraint string
}

// ParseRequirements
func ParseRequirements(str string) map[string]*semver.Constraints {
	requirements := map[string]*semver.Constraints{}
	// nolint: gomnd
	for _, req := range strings.Split(str, ",") {
		if req == "" {
			continue
		}
		splites := strings.SplitN(req, " ", 2)
		switch len(splites) {
		case 1:
			constraint, err := semver.NewConstraint("")
			if err != nil {
				continue
			}
			requirements[splites[0]] = constraint
		case 2:
			constraint, err := semver.NewConstraint(splites[1])
			if err != nil {
				continue
			}
			requirements[splites[0]] = constraint
		default:
			continue
		}
	}
	return requirements
}

func SemVersionBiggerThan(a, b string) bool {
	aver, err := semver.NewVersion(a)
	if err != nil {
		return false
	}
	bver, err := semver.NewVersion(b)
	if err != nil {
		return false
	}
	return aver.GreaterThan(bver)
}
