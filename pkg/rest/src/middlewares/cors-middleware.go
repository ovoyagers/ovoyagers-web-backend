package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{"http://localhost:5173", "https://ovoyagers.com", "https://www.ovoyagers.com", "http://localhost:4173", "http://localhost:3000", "https://bidmytour.com", "https://www.bidmytour.com"}

		origin := c.Request.Header.Get("Origin")
		allowOrigin := ""

		// Check if the origin is allowed
		for _, o := range allowedOrigins {
			if o == origin {
				allowOrigin = o
				break
			}
		}

		// If the origin is allowed, set the CORS headers
		if allowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			c.Writer.Header().Set("Vary", "Origin") // Helps with caching and CORS
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,  x-refresh-token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // Cache preflight for 1 day

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
