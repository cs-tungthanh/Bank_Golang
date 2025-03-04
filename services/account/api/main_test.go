package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// define Testmain to run the whole package tests
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
