package handler

import (
	// "fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	proxyURL := "https://registry-1.docker.io/"

	if strings.Contains(r.URL.Path,"registry-v2") {
		proxyURL = "https://production.cloudflare.docker.com/"
	}

	targetURL, err := url.Parse(proxyURL)
	if err != nil {
		http.Error(w, "目标服务器返回错误", 500)
		return
	}
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	// fmt.Fprintf(w, targetURL.String())
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = targetURL.Host
	proxy.ModifyResponse = func(r *http.Response) error {
		locationHeader := r.Header.Get("Location")
		if locationHeader != "" {
			url, _ := url.Parse(locationHeader)
			locationHeader = strings.Replace(locationHeader,url.Host, "docker.sonainai.com", -1)
			r.Header.Set("Location", locationHeader)
			log.Printf("Location: %s", locationHeader)
		}
		return nil
	}

	proxy.ServeHTTP(w, r)
}
