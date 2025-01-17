/*
Copyright 2025 Alexander Kharkevich

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
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	githubappjwt2tokenv1 "github.com/technicaldomain/github-app-jwt2token-controller/pkg/apis/githubappjwt2token/v1"
	scheme "github.com/technicaldomain/github-app-jwt2token-controller/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// GHSsGetter has a method to return a GHSInterface.
// A group's client should implement this interface.
type GHSsGetter interface {
	GHSs(namespace string) GHSInterface
}

// GHSInterface has methods to work with GHS resources.
type GHSInterface interface {
	Create(ctx context.Context, gHS *githubappjwt2tokenv1.GHS, opts metav1.CreateOptions) (*githubappjwt2tokenv1.GHS, error)
	Update(ctx context.Context, gHS *githubappjwt2tokenv1.GHS, opts metav1.UpdateOptions) (*githubappjwt2tokenv1.GHS, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, gHS *githubappjwt2tokenv1.GHS, opts metav1.UpdateOptions) (*githubappjwt2tokenv1.GHS, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*githubappjwt2tokenv1.GHS, error)
	List(ctx context.Context, opts metav1.ListOptions) (*githubappjwt2tokenv1.GHSList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *githubappjwt2tokenv1.GHS, err error)
	GHSExpansion
}

// gHSs implements GHSInterface
type gHSs struct {
	*gentype.ClientWithList[*githubappjwt2tokenv1.GHS, *githubappjwt2tokenv1.GHSList]
}

// newGHSs returns a GHSs
func newGHSs(c *GithubappV1Client, namespace string) *gHSs {
	return &gHSs{
		gentype.NewClientWithList[*githubappjwt2tokenv1.GHS, *githubappjwt2tokenv1.GHSList](
			"ghss",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *githubappjwt2tokenv1.GHS { return &githubappjwt2tokenv1.GHS{} },
			func() *githubappjwt2tokenv1.GHSList { return &githubappjwt2tokenv1.GHSList{} },
		),
	}
}
