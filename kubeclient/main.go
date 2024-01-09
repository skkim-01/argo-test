/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"syscall"

	// "flag"
	"fmt"
	// "path/filepath"
	"time"

	"skkim-01/github.com/argo-test/kubeclient/api"
	"skkim-01/github.com/argo-test/kubeclient/utils"

	JsonMapper "github.com/skkim-01/json-mapper/src"
	SignalWaiter "github.com/skkim-01/signal-waiter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	api.ThreadStart()

	SignalWaiter.Wait(
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGILL,
		syscall.SIGTRAP,
		syscall.SIGABRT,
		syscall.SIGSTKFLT,
		syscall.SIGBUS,
		syscall.SIGFPE,
		syscall.SIGSEGV)

	fmt.Println("done")
}

func bain() {
	var kubeconfig *string
	v := "/home/ubuntu/.kube/config"
	kubeconfig = &v
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	startTime := time.Now()
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(pods.Items[0])
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	endTime := time.Now()
	fmt.Println(`clientset.CoreV1().Pods("") Time taken:`, endTime.Sub(startTime))

	startTime = time.Now()
	yamlMap, err := JsonMapper.NewYamlObject(pods)
	if err != nil {
		panic(err.Error())
	}
	endTime = time.Now()
	utils.WriteFile("api.pods.yaml", yamlMap.Prints())
	fmt.Println("yaml result write Time taken:", endTime.Sub(startTime))

	startTime = time.Now()
	jsonMap, err := JsonMapper.NewObject(pods)
	if err != nil {
		panic(err.Error())
	}
	endTime = time.Now()
	utils.WriteFile("api.pods.json", jsonMap.PPrint())
	fmt.Println("json result write Time taken:", endTime.Sub(startTime))

	fmt.Println()
	// test console
	nItems := len(jsonMap.Find("items").([]interface{}))
	listPodInfo := make([]map[string]string, nItems)
	for i := 0; i < nItems; i++ {
		listPodInfo[i] = make(map[string]string)
		listPodInfo[i]["name"] = jsonMap.Find(fmt.Sprintf("items.%v.metadata.name", i)).(string)
		listPodInfo[i]["namespace"] = jsonMap.Find(fmt.Sprintf("items.%v.metadata.namespace", i)).(string)
		tmp := jsonMap.Find(fmt.Sprintf("items.%v.spec.nodeName", i))
		if tmp != nil {
			listPodInfo[i]["node"] = tmp.(string)
		} else {
			listPodInfo[i]["node"] = ""
		}
		listPodInfo[i]["status"] = jsonMap.Find(fmt.Sprintf("items.%v.status.phase", i)).(string)

		//fmt.Println(jsonMap.Find(fmt.Sprintf("items.%v.metadata.name", i)))
		//fmt.Println(jsonMap.Find(fmt.Sprintf("items.%v.metadata.namespace", i)))
		//fmt.Println(jsonMap.Find(fmt.Sprintf("items.%v.status.phase", i)))
		//fmt.Println(jsonMap.Find(fmt.Sprintf("items.%v.spec.nodeName", i)))
	}

	fmt.Println(JsonMapper.Convert(listPodInfo))
	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := "default"
	// pod := "example-xxxxx"
	// _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 		pod, namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// }

	//time.Sleep(10 * time.Second)
}
