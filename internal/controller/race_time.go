package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTotalRaceTime(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
