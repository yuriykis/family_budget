package apiserver

import (
	"app/internal/app/controllers"
	"app/internal/app/controllers/controller"
	"app/internal/app/store"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router            *gin.Engine
	logger            *log.Logger
	userController    controllers.IUserController
	serviceController controllers.IServiceController
}

func newServer(store store.IStore) *server {
	s := &server{
		router:            gin.Default(),
		logger:            log.New(),
		userController:    controller.NewsUserController(store),
		serviceController: controller.NewServiceController(),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.GET("/api/healthcheck", s.handleHealthcheck())
	s.router.POST("/api/register", s.handleRegister())
	s.router.POST("/api/auth", s.handleAuth())
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}

func (s *server) handleHealthcheck() gin.HandlerFunc {
	return s.serviceController.CheckHealth
}

func (s *server) handleRegister() gin.HandlerFunc {
	return s.userController.RegisterUser
}

func (s *server) handleAuth() gin.HandlerFunc {
	return s.userController.AthenticateUser
}
