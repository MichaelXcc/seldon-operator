/*
Copyright 2019 The Seldon Team.

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

package v1alpha2

import (
	"fmt"
	"k8s.io/api/core/v1"
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestCleanImageName(t *testing.T) {
	//g := gomega.NewGomegaWithT(t)
	name2 := cleanContainerName("AB_C")
	if name2 != "ab-c" {
		t.Fatalf("should be abc: %s", name2)
	}
}

func TestCleanDeploymentName(t *testing.T) {
	mlDep := &SeldonDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mymodel",
		},
		Spec: SeldonDeploymentSpec{
			Name: "mymodel",
			Predictors: []PredictorSpec{
				{
					Name: "mymodel",
					ComponentSpecs: []*SeldonPodSpec{
						&SeldonPodSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name:  "classifier",
										Image: "seldonio/mock_classifier:1.0",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	name := GetDeploymentName(mlDep, mlDep.Spec.Predictors[0], mlDep.Spec.Predictors[0].ComponentSpecs[0])
	fmt.Println(name)
}

func TestStorageSeldonDeployment(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &SeldonDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		}}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &SeldonDeployment{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}
