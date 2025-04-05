package handler

import (
	// "fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // 设置跨域头部信息
    w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名访问
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

    // 处理预检请求（OPTIONS）
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // 解析目标 URL
    targetURL, err := url.Parse("https://surrit.com")
    if err != nil {
        http.Error(w, "目标服务器返回错误", 500)
        return
    }

    // 创建反向代理
    proxy := httputil.NewSingleHostReverseProxy(targetURL)
    r.URL.Host = targetURL.Host
    r.URL.Scheme = targetURL.Scheme
    r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
    r.Host = targetURL.Host

    // 调用反向代理处理请求
    proxy.ServeHTTP(w, r)
}
