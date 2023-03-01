package apiserver

import (
	"categoryservice/internal/app/handlers"
	"categoryservice/internal/app/handlers/handler"
	"categoryservice/internal/app/middlewares"
	"categoryservice/internal/app/store"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router          *gin.Engine
	logger          *log.Logger
	serviceHandler  handlers.IServiceHandler
	categoryHandler handlers.ICategoryHandler
}

func newServer(store store.IStore) *server {
	s := &server{
		router:          gin.Default(),
		logger:          log.New(),
		serviceHandler:  handler.NewServiceHandler(),
		categoryHandler: handler.NewCategoryHandler(store),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	s.router.Use(middlewares.CorsMiddleware())

	public := s.router.Group("/api")
	public.GET("/healthcheck", s.serviceHandler.CheckHealth)

	protected := s.router.Group("/api/auth")
	protected.Use(middlewares.JwtMiddleware())

	protected.POST("/category", s.categoryHandler.CreateCategory)
	protected.GET("/category", s.categoryHandler.FindAllCategories)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
