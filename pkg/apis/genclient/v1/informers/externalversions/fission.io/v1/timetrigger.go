/*
Copyright The Fission Authors.

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
	time "time"

	fissioniov1 "github.com/fission/fission/pkg/apis/fission.io/v1"
	versioned "github.com/fission/fission/pkg/apis/genclient/v1/clientset/versioned"
	internalinterfaces "github.com/fission/fission/pkg/apis/genclient/v1/informers/externalversions/internalinterfaces"
	v1 "github.com/fission/fission/pkg/apis/genclient/v1/listers/fission.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TimeTriggerInformer provides access to a shared informer and lister for
// TimeTriggers.
type TimeTriggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TimeTriggerLister
}

type _timeTriggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTimeTriggerInformer constructs a new informer for TimeTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTimeTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTimeTriggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTimeTriggerInformer constructs a new informer for TimeTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTimeTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.V1V1().TimeTriggers(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.V1V1().TimeTriggers(namespace).Watch(options)
			},
		},
		&fissioniov1.TimeTrigger{},
		resyncPeriod,
		indexers,
	)
}

func (f *_timeTriggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTimeTriggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *_timeTriggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fissioniov1.TimeTrigger{}, f.defaultInformer)
}

func (f *_timeTriggerInformer) Lister() v1.TimeTriggerLister {
	return v1.NewTimeTriggerLister(f.Informer().GetIndexer())
}