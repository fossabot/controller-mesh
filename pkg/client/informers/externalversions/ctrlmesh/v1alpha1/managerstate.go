/*
Copyright 2023 The KusionStack Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	ctrlmeshv1alpha1 "github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh/v1alpha1"
	versioned "github.com/KusionStack/controller-mesh/pkg/client/clientset/versioned"
	internalinterfaces "github.com/KusionStack/controller-mesh/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/KusionStack/controller-mesh/pkg/client/listers/ctrlmesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ManagerStateInformer provides access to a shared informer and lister for
// ManagerStates.
type ManagerStateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ManagerStateLister
}

type managerStateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewManagerStateInformer constructs a new informer for ManagerState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagerStateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagerStateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredManagerStateInformer constructs a new informer for ManagerState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagerStateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CtrlmeshV1alpha1().ManagerStates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CtrlmeshV1alpha1().ManagerStates().Watch(context.TODO(), options)
			},
		},
		&ctrlmeshv1alpha1.ManagerState{},
		resyncPeriod,
		indexers,
	)
}

func (f *managerStateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagerStateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managerStateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ctrlmeshv1alpha1.ManagerState{}, f.defaultInformer)
}

func (f *managerStateInformer) Lister() v1alpha1.ManagerStateLister {
	return v1alpha1.NewManagerStateLister(f.Informer().GetIndexer())
}
