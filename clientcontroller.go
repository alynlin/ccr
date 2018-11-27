package main

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	samplecrdv1 "ccr/pkg/apis/samplecrd/v1"
	"ccr/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/watch"
	"github.com/golang/glog"
)

func watchNetwork(networkClient *versioned.Clientset) cache.SharedIndexInformer {
	indexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return networkClient.SamplecrdV1().Networks("").List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {

				return networkClient.SamplecrdV1().Networks("").Watch(options)
			},
		},
		&samplecrdv1.Network{},
		0,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	indexInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    addNetwor,
			UpdateFunc: updateNetwor,
			DeleteFunc: deleteNetwor,
		})
	return indexInformer
}

func addNetwor(obj interface{}) {
	glog.Info("================addNetwork================:%s", "add")
}

func updateNetwor(oldObj, obj interface{}) {
	glog.Info("================updateNetwork================:%s", "update")
}

func deleteNetwor(obj interface{}) {
	glog.Info("================deleteNetwork================:%s", "delete")
}
