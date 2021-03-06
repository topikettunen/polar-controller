/*
Copyright 2020 Topi Kettunen.

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

	polarsquadv1 "github.com/topikettunen/polar-controller/pkg/apis/polarsquad/v1"
	versioned "github.com/topikettunen/polar-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/topikettunen/polar-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/topikettunen/polar-controller/pkg/client/listers/polarsquad/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PolarDeploymentInformer provides access to a shared informer and lister for
// PolarDeployments.
type PolarDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PolarDeploymentLister
}

type polarDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPolarDeploymentInformer constructs a new informer for PolarDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPolarDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPolarDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPolarDeploymentInformer constructs a new informer for PolarDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPolarDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolarsquadV1().PolarDeployments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolarsquadV1().PolarDeployments(namespace).Watch(context.TODO(), options)
			},
		},
		&polarsquadv1.PolarDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *polarDeploymentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPolarDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *polarDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&polarsquadv1.PolarDeployment{}, f.defaultInformer)
}

func (f *polarDeploymentInformer) Lister() v1.PolarDeploymentLister {
	return v1.NewPolarDeploymentLister(f.Informer().GetIndexer())
}
