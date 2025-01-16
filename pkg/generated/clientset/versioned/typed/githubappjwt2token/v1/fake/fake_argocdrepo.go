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

package fake

import (
	v1 "github.com/kharkevich/github-app-jwt2token-controller/pkg/apis/githubappjwt2token/v1"
	githubappjwt2tokenv1 "github.com/kharkevich/github-app-jwt2token-controller/pkg/generated/clientset/versioned/typed/githubappjwt2token/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeArgoCDRepos implements ArgoCDRepoInterface
type fakeArgoCDRepos struct {
	*gentype.FakeClientWithList[*v1.ArgoCDRepo, *v1.ArgoCDRepoList]
	Fake *FakeGithubappV1
}

func newFakeArgoCDRepos(fake *FakeGithubappV1, namespace string) githubappjwt2tokenv1.ArgoCDRepoInterface {
	return &fakeArgoCDRepos{
		gentype.NewFakeClientWithList[*v1.ArgoCDRepo, *v1.ArgoCDRepoList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("argocdrepos"),
			v1.SchemeGroupVersion.WithKind("ArgoCDRepo"),
			func() *v1.ArgoCDRepo { return &v1.ArgoCDRepo{} },
			func() *v1.ArgoCDRepoList { return &v1.ArgoCDRepoList{} },
			func(dst, src *v1.ArgoCDRepoList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ArgoCDRepoList) []*v1.ArgoCDRepo { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ArgoCDRepoList, items []*v1.ArgoCDRepo) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
