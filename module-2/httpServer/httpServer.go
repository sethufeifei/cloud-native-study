package main

import (
	"flag"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	flag.Set("V", "4")

	http.HandleFunc("/healthz", healthz)

	glog.V(2).Info("Starting http server...")
	listen(":80")
	glog.V(2).Info("http server is started")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
	w.WriteHeader(http.StatusOK)
}

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	glog.V(2).Info("access_url = ", r.RequestURI, ", client_host = ", r.Host)

	addVersion(w, r)
	addHeader(w, r)

}

func addHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	glog.V(2).Info("headerInterceptor, header = ", header)
	for headerKey, headerValueList := range header {
		for _, headerValue := range headerValueList {
			w.Header().Set(headerKey, headerValue)
		}
	}
}

func addVersion(w http.ResponseWriter, r *http.Request) {
	//os.Setenv("VERSION", "1.0.0.1")
	version := os.Getenv("VERSION")
	glog.V(2).Info("versionInterceptor, version = ", version)
	w.Header().Add("VERSION", version)
}

func listen(addr string) {
	var server httpServer
	err := http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatal(err)
	}
}
