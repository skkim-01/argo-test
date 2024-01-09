package cli

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"skkim-01/github.com/argo-test/kubeclient/utils"

	JsonMapper "github.com/skkim-01/json-mapper/src"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Pods(w http.ResponseWriter, r *http.Request) {
	lt := utils.GetLocalTime()
	st := time.Now()

	// var kubeconfig *string
	// v := "/home/ubuntu/.kube/config"
	// kubeconfig = &v

	// // use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // create the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// call k8s api
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// convert json
	jsonMap, err := JsonMapper.NewObject(pods)
	if err != nil {
		panic(err.Error())
	}
	// make result
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
	}

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// return
	fmt.Fprintf(w, `%v`, JsonMapper.Convert(listPodInfo))
	et := time.Now()

	fmt.Println(lt, "\033[31m", "\t [GET] /api/v1/cli/pods \t", "\033[0m", et.Sub(st))
}
