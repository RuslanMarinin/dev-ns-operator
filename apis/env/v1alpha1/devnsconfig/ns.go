/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package devnsconfig

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	envv1alpha1 "github.com/RuslanMarinin/dev-ns-operator/apis/env/v1alpha1"
	"github.com/RuslanMarinin/dev-ns-operator/apis/env/v1alpha1/devnsconfig/mutate"
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespaceName creates the Namespace resource with name parent.Spec.Name.
func CreateNamespaceName(
	parent *envv1alpha1.DevNS,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": parent.Spec.Name, //  controlled by field: Name
				"labels": map[string]interface{}{
					"name":       parent.Spec.Name,       //  controlled by field: Name
					"jiraTicket": parent.Spec.JiraTicket, //  controlled by field: jiraTicket
					"gitBranch":  parent.Spec.GitBranch,  //  controlled by field: gitBranch
				},
			},
		},
	}

	return mutate.MutateNamespaceName(resourceObj, parent, reconciler, req)
}
