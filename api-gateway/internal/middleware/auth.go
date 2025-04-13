package middleware

import (
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // In a real application, you would validate a JWT or API key here
        apiKey := c.GetHeader("X-API-Key")
        if apiKey == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        // For this example, we'll accept any non-empty API key
        c.Next()
    }
}

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Log request details
        log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
    }
}