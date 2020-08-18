/*
Copyright 2018 The Kong Authors.

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
	"context"
	time "time"

	configurationv1 "github.com/kong/kubernetes-ingress-controller/pkg/apis/configuration/v1"
	versioned "github.com/kong/kubernetes-ingress-controller/pkg/client/configuration/clientset/versioned"
	internalinterfaces "github.com/kong/kubernetes-ingress-controller/pkg/client/configuration/informers/externalversions/internalinterfaces"
	v1 "github.com/kong/kubernetes-ingress-controller/pkg/client/configuration/listers/configuration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KongClusterPluginInformer provides access to a shared informer and lister for
// KongClusterPlugins.
type KongClusterPluginInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KongClusterPluginLister
}

type kongClusterPluginInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKongClusterPluginInformer constructs a new informer for KongClusterPlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKongClusterPluginInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKongClusterPluginInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKongClusterPluginInformer constructs a new informer for KongClusterPlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKongClusterPluginInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigurationV1().KongClusterPlugins().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigurationV1().KongClusterPlugins().Watch(context.TODO(), options)
			},
		},
		&configurationv1.KongClusterPlugin{},
		resyncPeriod,
		indexers,
	)
}

func (f *kongClusterPluginInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKongClusterPluginInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kongClusterPluginInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configurationv1.KongClusterPlugin{}, f.defaultInformer)
}

func (f *kongClusterPluginInformer) Lister() v1.KongClusterPluginLister {
	return v1.NewKongClusterPluginLister(f.Informer().GetIndexer())
}