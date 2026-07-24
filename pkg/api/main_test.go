package api

import (
	"testing"

	"github.com/hanzo-fi/go-libs/v3/logging"
	"github.com/hanzo-fi/go-libs/v3/testing/docker"
	"github.com/hanzo-fi/go-libs/v3/testing/utils"

	"github.com/hanzo-fi/go-libs/v3/testing/platform/pgtesting"
)

var srv *pgtesting.PostgresServer

func TestMain(m *testing.M) {
	utils.WithTestMain(func(t *utils.TestingTForMain) int {
		srv = pgtesting.CreatePostgresServer(t, docker.NewPool(t, logging.Testing()))

		return m.Run()
	})
}
