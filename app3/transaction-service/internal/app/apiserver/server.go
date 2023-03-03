package apiserver

import (
	"transactionservice/internal/app/handlers"
	"transactionservice/internal/app/handlers/handler"
	"transactionservice/internal/app/middlewares"
	"transactionservice/internal/app/store"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router               *gin.Engine
	logger               *log.Logger
	serviceHealthHandler handlers.IServiceHealthHandler
	transactionHandler   handlers.ITransactionHandler
}

func newServer(store store.IStore) *server {
	s := &server{
		router:               gin.Default(),
		logger:               log.New(),
		serviceHealthHandler: handler.NewServiceHealthHandler(),
		transactionHandler:   handler.NewTransactionHandler(store),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	s.router.Use(middlewares.CorsMiddleware())

	public := s.router.Group("/api")
	public.GET("/healthcheck", s.serviceHealthHandler.CheckHealth)

	protected := s.router.Group("/api/auth")
	protected.Use(middlewares.JwtMiddleware())

	protected.POST("/budget/:budget_id/transaction", s.transactionHandler.CreateTransaction)
	protected.GET("/budget/:budget_id/transaction", s.transactionHandler.GetTransactions)

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
