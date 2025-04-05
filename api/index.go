package handler

import (
	// "fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	targetURL, err := url.Parse("https://surrit.com")
	if err != nil {
		http.Error(w, "目标服务器返回错误", 500)
		return
	}

	// fmt.Fprintf(w, targetURL.String())
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = targetURL.Host
	proxy.ServeHTTP(w, r)
}
