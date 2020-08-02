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

package v1

import (
	"context"
	jackistomv1 "github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom/v1"
	versioned "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jackistom/kubectl-Ecs/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jackistom/kubectl-Ecs/pkg/client/listers/jackistom/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SylixosInformer provides access to a shared informer and lister for
// Sylixoses.
type SylixosInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SylixosLister
}

type sylixosInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSylixosInformer constructs a new informer for Sylixos type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSylixosInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSylixosInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSylixosInformer constructs a new informer for Sylixos type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSylixosInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JackistomV1().Sylixoses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JackistomV1().Sylixoses(namespace).Watch(context.TODO(), options)
			},
		},
		&jackistomv1.Sylixos{},
		resyncPeriod,
		indexers,
	)
}

func (f *sylixosInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSylixosInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sylixosInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jackistomv1.Sylixos{}, f.defaultInformer)
}

func (f *sylixosInformer) Lister() v1.SylixosLister {
	return v1.NewSylixosLister(f.Informer().GetIndexer())
}
