package apiserver

import (
	"app/internal/app/store"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router *gin.Engine
	logger *log.Logger
	store  store.IStore
}

func newServer(store store.IStore) *server {
	s := &server{
		router: gin.Default(),
		logger: log.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.GET("/ping", s.handlePing())
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "pong")
	}
}
