package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
)

const (
	maxRequests     = 100
	perMinutePeriod = time.Minute * 1
)

var (
	ipRequestsCounts = make(map[string]int)
	mutex            = &sync.Mutex{}
)

func RateLimiterMiddleware(c *gin.Context) {
	ip := c.ClientIP()

	mutex.Lock()
	defer mutex.Unlock()

	count := ipRequestsCounts[ip]
	if count >= maxRequests {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, models.Error{
			Message:    "too many requests",
			Error:      "Status Too Many Requests",
			Status:     "error",
			StatusCode: http.StatusTooManyRequests,
		})
		return
	}

	ipRequestsCounts[ip] = count + 1
	time.AfterFunc(perMinutePeriod, func() {
		mutex.Lock()
		defer mutex.Unlock()

		ipRequestsCounts[ip] = ipRequestsCounts[ip] - 1
	})
	c.Next()
}
