package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	isActive      = true
	isActiveMutex sync.Mutex
)

type response struct {
	Status string `json:"status"`
}

func HealthCheck(c *gin.Context) {
	isActiveMutex.Lock()
	defer isActiveMutex.Unlock()

	c.Header("no-gzip", "health-check-disabled-gzip")
	c.Header("cache-control", "no-cache")
	c.JSON(http.StatusOK, response{"OK"})

}
