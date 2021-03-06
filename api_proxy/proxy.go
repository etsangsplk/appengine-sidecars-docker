// Package main starts an App Engine API proxy on port 10001.
// It is intended to allow migration of APIs from the Service Bridge to
// publically available APIs and allow users & runtimes to migrate at a
// later date.
package main // import "github.com/GoogleCloudPlatform/appengine-sidecars-docker/api_proxy"

import (
	"log"
	"net/http"
)

func sbEnabledError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("App Engine APIs over the Service Bridge are disabled.\n" +
		"If they are required, enable them by setting the following to your app.yaml:\n" +
		"\n" +
		"beta_settings:\n" +
		"  enable_app_engine_apis: true\n"))
}

func makeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sbEnabledError)
	return mux
}

func main() {
	handler := makeHandler()
	s := &http.Server{
		Addr:    ":10001",
		Handler: handler,
	}
	log.Fatal(s.ListenAndServe())
}
