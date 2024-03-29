package apiserver

import (
	"userservice/internal/app/handlers"
	"userservice/internal/app/handlers/handler"
	"userservice/internal/app/middlewares"
	"userservice/internal/app/store"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router               *gin.Engine
	logger               *log.Logger
	userHandler          handlers.IUserHandler
	serviceHealthHandler handlers.IServiceHealthHandler
}

func newServer(store store.IStore) *server {
	s := &server{
		router:               gin.Default(),
		logger:               log.New(),
		userHandler:          handler.NewsUserHandler(store),
		serviceHealthHandler: handler.NewServiceHealthHandler(),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	s.router.Use(middlewares.CorsMiddleware())

	public := s.router.Group("/api")
	public.GET("/healthcheck", s.serviceHealthHandler.CheckHealth)
	public.POST("/register", s.userHandler.RegisterUser)
	public.POST("/login", s.userHandler.AthenticateUser)

	protected := s.router.Group("/api/auth")
	protected.Use(middlewares.JwtMiddleware())
	protected.GET("/user", s.userHandler.GetUser)
	protected.GET("/user/:user_id", s.userHandler.GetUserByID)

	protected.PUT("/user/:user_id", s.userHandler.UpdateUser)
	protected.DELETE("/user/:user_id", s.userHandler.DeleteUser)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
