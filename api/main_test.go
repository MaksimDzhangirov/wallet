package api

import (
	db "github.com/MaksimDzhangirov/wallet/db/sqlc"
	"github.com/MaksimDzhangirov/wallet/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		WithdrawalLimit: 2,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}