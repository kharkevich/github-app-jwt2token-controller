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
	v1 "github.com/technicaldomain/github-app-jwt2token-controller/pkg/apis/githubappjwt2token/v1"
	githubappjwt2tokenv1 "github.com/technicaldomain/github-app-jwt2token-controller/pkg/generated/clientset/versioned/typed/githubappjwt2token/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeDockerConfigJsons implements DockerConfigJsonInterface
type fakeDockerConfigJsons struct {
	*gentype.FakeClientWithList[*v1.DockerConfigJson, *v1.DockerConfigJsonList]
	Fake *FakeGithubappV1
}

func newFakeDockerConfigJsons(fake *FakeGithubappV1, namespace string) githubappjwt2tokenv1.DockerConfigJsonInterface {
	return &fakeDockerConfigJsons{
		gentype.NewFakeClientWithList[*v1.DockerConfigJson, *v1.DockerConfigJsonList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("dockerconfigjsons"),
			v1.SchemeGroupVersion.WithKind("DockerConfigJson"),
			func() *v1.DockerConfigJson { return &v1.DockerConfigJson{} },
			func() *v1.DockerConfigJsonList { return &v1.DockerConfigJsonList{} },
			func(dst, src *v1.DockerConfigJsonList) { dst.ListMeta = src.ListMeta },
			func(list *v1.DockerConfigJsonList) []*v1.DockerConfigJson { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.DockerConfigJsonList, items []*v1.DockerConfigJson) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
