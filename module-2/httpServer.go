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
	flag.Set("V", "4")

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/header", headerInterceptor(header))
	http.HandleFunc("/version", versionInterceptor(version))

	glog.V(2).Info("Starting http server...")
	listen(":80")
	glog.V(2).Info("http server is started")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func header(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func version(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func listen(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func headerInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		glog.V(2).Info("headerInterceptor, header = ", header)
		for headerKey, headerValueList := range header {
			for _, headerValue := range headerValueList {
				w.Header().Set(headerKey, headerValue)
			}
		}
		h(w, r)
	}
}

func versionInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//os.Setenv("VERSION", "1.0.0.1")
		version := os.Getenv("VERSION")
		glog.V(2).Info("versionInterceptor, version = ", version)
		w.Header().Add("VERSION", version)
		h(w, r)
	}
}
