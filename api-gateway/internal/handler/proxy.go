package handler

import (
    "net/http"
    "net/http/httputil"
    "net/url"

    "github.com/gin-gonic/gin"
)

func NewInventoryServiceProxy(target string) gin.HandlerFunc {
    return func(c *gin.Context) {
        remote, err := url.Parse(target)
        if err != nil {
            panic(err)
        }

        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = c.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.URL.Path = "/api" + c.Param("path")
        }

        proxy.ServeHTTP(c.Writer, c.Request)
    }
}

func NewOrderServiceProxy(target string) gin.HandlerFunc {
    return func(c *gin.Context) {
        remote, err := url.Parse(target)
        if err != nil {
            panic(err)
        }

        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = c.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.URL.Path = "/api" + c.Param("path")
        }

        proxy.ServeHTTP(c.Writer, c.Request)
    }
}