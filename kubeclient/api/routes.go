package api

import (
	"fmt"
	"net/http"

	v1 "skkim-01/github.com/argo-test/kubeclient/api/v1"
	"skkim-01/github.com/argo-test/kubeclient/api/v1/cli"
)

func ThreadStart() {
	go start()
}

func start() {
	http.HandleFunc("/api/v1/cli/pods", cli.Pods)
	http.HandleFunc("/api/v1/stat", v1.Stat)

	// serve file
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("#Webserver Thread \t start with 9999")
	http.ListenAndServe(fmt.Sprintf(":%v", 9999), nil)
}
