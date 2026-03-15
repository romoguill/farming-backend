package handler

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	router = gin.Default()

	os.Exit(m.Run())
}
