package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	requests      int           // allowed requests per window
	window        time.Duration // window size
	lastResetTime time.Time
	count         int
	mu            sync.Mutex
}

var limiter = rateLimiter{
	requests:      5,
	window:        10 * time.Second,
	lastResetTime: time.Now(),
}

// RequestLogger logs basic information about each request.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)

		log.Printf("%s %s -> %d (%s)", method, path, status, duration)
	}
}

// RateLimite
func Ratelimitter() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		//Reset window
		if time.Since(limiter.lastResetTime) > limiter.window {
			limiter.count = 0
			limiter.lastResetTime = time.Now()
		}

		//Check if requests exceed limit
		if limiter.count >= limiter.requests {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		limiter.count++
		c.Next()
	}
}

// CORSMiddleware configures and returns a CORS middleware handler.
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	if len(allowedOrigins) == 0 {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = allowedOrigins
	}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}

	return cors.New(config)
}
