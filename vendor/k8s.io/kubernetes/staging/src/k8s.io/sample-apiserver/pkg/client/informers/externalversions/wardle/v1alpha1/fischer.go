/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	wardle_v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
	versioned "k8s.io/sample-apiserver/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/sample-apiserver/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/sample-apiserver/pkg/client/listers/wardle/v1alpha1"
	time "time"
)

// FischerInformer provides access to a shared informer and lister for
// Fischers.
type FischerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FischerLister
}

type fischerInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewFischerInformer constructs a new informer for Fischer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFischerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.WardleV1alpha1().Fischers().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.WardleV1alpha1().Fischers().Watch(options)
			},
		},
		&wardle_v1alpha1.Fischer{},
		resyncPeriod,
		indexers,
	)
}

func defaultFischerInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFischerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *fischerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wardle_v1alpha1.Fischer{}, defaultFischerInformer)
}

func (f *fischerInformer) Lister() v1alpha1.FischerLister {
	return v1alpha1.NewFischerLister(f.Informer().GetIndexer())
}
