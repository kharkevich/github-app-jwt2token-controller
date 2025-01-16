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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apisgithubappjwt2tokenv1 "github.com/kharkevich/github-app-jwt2token-controller/pkg/apis/githubappjwt2token/v1"
	versioned "github.com/kharkevich/github-app-jwt2token-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/kharkevich/github-app-jwt2token-controller/pkg/generated/informers/externalversions/internalinterfaces"
	githubappjwt2tokenv1 "github.com/kharkevich/github-app-jwt2token-controller/pkg/generated/listers/githubappjwt2token/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GHSInformer provides access to a shared informer and lister for
// GHSs.
type GHSInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() githubappjwt2tokenv1.GHSLister
}

type gHSInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGHSInformer constructs a new informer for GHS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGHSInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGHSInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGHSInformer constructs a new informer for GHS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGHSInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GithubappV1().GHSs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GithubappV1().GHSs(namespace).Watch(context.TODO(), options)
			},
		},
		&apisgithubappjwt2tokenv1.GHS{},
		resyncPeriod,
		indexers,
	)
}

func (f *gHSInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGHSInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gHSInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisgithubappjwt2tokenv1.GHS{}, f.defaultInformer)
}

func (f *gHSInformer) Lister() githubappjwt2tokenv1.GHSLister {
	return githubappjwt2tokenv1.NewGHSLister(f.Informer().GetIndexer())
}
