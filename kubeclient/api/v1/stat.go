package v1

import (
	"fmt"
	"net/http"
	"time"

	"skkim-01/github.com/argo-test/kubeclient/utils"
)

func Stat(w http.ResponseWriter, r *http.Request) {
	lt := utils.GetLocalTime()
	st := time.Now()

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	fmt.Fprintf(w, `{"status": 200 Ok}`)

	et := time.Now()
	fmt.Println(lt, "\033[34m", "\t [GET] /api/v1/stat \t", "\033[0m", et.Sub(st))
}
