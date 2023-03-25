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
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	envv1alpha1 "github.com/RuslanMarinin/dev-ns-operator/apis/env/v1alpha1"
)

// sampleDevNS is a sample containing all fields
const sampleDevNS = `apiVersion: env.ruslan.marinin/v1alpha1
kind: DevNS
metadata:
  name: devns-sample
spec:
  Name: "development"
  jiraTicket: "JIRA-001"
  gitBranch: "feature/JIRA-001"
`

// sampleDevNSRequired is a sample containing only required fields
const sampleDevNSRequired = `apiVersion: env.ruslan.marinin/v1alpha1
kind: DevNS
metadata:
  name: devns-sample
spec:
  Name: "development"
  jiraTicket: "JIRA-001"
  gitBranch: "feature/JIRA-001"
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleDevNSRequired
	}

	return sampleDevNS
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj envv1alpha1.DevNS,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*envv1alpha1.DevNS,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateNamespaceName,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*envv1alpha1.DevNS,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){}

func ConvertWorkload(component workload.Workload) (*envv1alpha1.DevNS, error) {
	p, ok := component.(*envv1alpha1.DevNS)
	if !ok {
		return nil, envv1alpha1.ErrUnableToConvertDevNS
	}

	return p, nil
}
