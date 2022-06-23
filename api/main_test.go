package api

import (
	"os"
	"testing"
	"time"

	db "github.com/cs-tungthanh/Bank_Golang/db/sqlc"
	"github.com/cs-tungthanh/Bank_Golang/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

// define Testmain to run the whole package tests
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
