package main

import (
	"GO_thinknetica/01-lesson/demoapp/pkg/strings_utils"
	"net/http"
)

func main() {
	http.ListenAndServe(
		":8181",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				s := r.URL.Query()["s"]
				rev := strings_utils.Rev(s)
				w.Write([]byte(rev))
			},
		),
	)
}
