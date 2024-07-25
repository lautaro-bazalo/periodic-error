package routes

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, response{"OK"})
}
