/*
Copyright The Kubernetes Authors.

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

package v2

import (
	bk_bcs_v2 "bk-bcs/bcs-services/bcs-log-webhook-server/pkg/apis/bk-bcs/v2"
	versioned "bk-bcs/bcs-services/bcs-log-webhook-server/pkg/client/clientset/versioned"
	internalinterfaces "bk-bcs/bcs-services/bcs-log-webhook-server/pkg/client/informers/externalversions/internalinterfaces"
	v2 "bk-bcs/bcs-services/bcs-log-webhook-server/pkg/client/listers/bk-bcs/v2"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BcsLogConfigInformer provides access to a shared informer and lister for
// BcsLogConfigs.
type BcsLogConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.BcsLogConfigLister
}

type bcsLogConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBcsLogConfigInformer constructs a new informer for BcsLogConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBcsLogConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBcsLogConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBcsLogConfigInformer constructs a new informer for BcsLogConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBcsLogConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV2().BcsLogConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV2().BcsLogConfigs(namespace).Watch(options)
			},
		},
		&bk_bcs_v2.BcsLogConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *bcsLogConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBcsLogConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bcsLogConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bk_bcs_v2.BcsLogConfig{}, f.defaultInformer)
}

func (f *bcsLogConfigInformer) Lister() v2.BcsLogConfigLister {
	return v2.NewBcsLogConfigLister(f.Informer().GetIndexer())
}