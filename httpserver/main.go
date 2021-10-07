package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type ResponseWriter interface {
	http.ResponseWriter
	Status() int
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

var _ ResponseWriter = &responseWriter{}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.status = statusCode
}

// rootHandler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	requestHeader := r.Header
	for k, v := range requestHeader{
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		for _, s := range v{
			w.Header().Add(k, s)
		}
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	clientIp := getClientIp(r)
	//fmt.Println(w.Header())
	fmt.Printf("client ip: %s, http response code: %s\n", clientIp, w.Header().Get("Status Code"))
}

// healthzHandler
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// getClientIp
func getClientIp(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		//测试用，实际场景不能直接返回
		return forwarded
	}
	if realIp := r.Header.Get("X-Real-IP"); realIp != "" {
		//测试用，实际场景不能直接返回
		return realIp
	}

	remoteIp, _, _ := net.SplitHostPort(r.RemoteAddr)

	return remoteIp
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
