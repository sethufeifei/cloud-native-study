package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	flag.Set("V", "4")

	//NewServeMux可以创建一个ServeMux实例，ServeMux同时也实现了ServeHTTP方法，因此代码中的mux也是一种handler。
	//把它当成参数传给http.ListenAndServe方法，后者会把mux传给Server实例。
	//因为指定了handler，因此整个http服务就不再是DefaultServeMux，而是mux，无论是在注册路由还是提供请求服务的时候
	mux := http.NewServeMux()
	// 06. debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/", index)

	glog.V(2).Info("Starting http server...")
	listen(":80", mux)
	glog.V(2).Info("http server is started")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "ok")
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("access_url = %s, client_host = %s", r.RequestURI, ClientIP(r))

	addVersion(w, r)
	addHeader(w, r)

	w.WriteHeader(http.StatusOK)
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
//解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
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

func listen(addr string, mux http.Handler) {
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
