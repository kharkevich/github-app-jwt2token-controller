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
	http "net/http"

	githubappjwt2tokenv1 "github.com/technicaldomain/github-app-jwt2token-controller/pkg/apis/githubappjwt2token/v1"
	scheme "github.com/technicaldomain/github-app-jwt2token-controller/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type GithubappV1Interface interface {
	RESTClient() rest.Interface
	ArgoCDReposGetter
	DockerConfigJsonsGetter
	GHSsGetter
}

// GithubappV1Client is used to interact with features provided by the githubapp.technicaldomain.xyz group.
type GithubappV1Client struct {
	restClient rest.Interface
}

func (c *GithubappV1Client) ArgoCDRepos(namespace string) ArgoCDRepoInterface {
	return newArgoCDRepos(c, namespace)
}

func (c *GithubappV1Client) DockerConfigJsons(namespace string) DockerConfigJsonInterface {
	return newDockerConfigJsons(c, namespace)
}

func (c *GithubappV1Client) GHSs(namespace string) GHSInterface {
	return newGHSs(c, namespace)
}

// NewForConfig creates a new GithubappV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*GithubappV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new GithubappV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*GithubappV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &GithubappV1Client{client}, nil
}

// NewForConfigOrDie creates a new GithubappV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GithubappV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GithubappV1Client for the given RESTClient.
func New(c rest.Interface) *GithubappV1Client {
	return &GithubappV1Client{c}
}

func setConfigDefaults(config *rest.Config) {
	gv := githubappjwt2tokenv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GithubappV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}