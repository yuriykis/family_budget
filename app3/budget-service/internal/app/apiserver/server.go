package apiserver

import (
	"budgetservice/internal/app/handlers"
	"budgetservice/internal/app/handlers/handler"
	"budgetservice/internal/app/middlewares"
	"budgetservice/internal/app/store"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router               *gin.Engine
	logger               *log.Logger
	serviceHealthHandler handlers.IServiceHealthHandler
	budgetHandler        handlers.IBudgetHandler
}

func newServer(store store.IStore) *server {
	s := &server{
		router:               gin.Default(),
		logger:               log.New(),
		serviceHealthHandler: handler.NewServiceHealthHandler(),
		budgetHandler:        handler.NewBudgetHandler(store),
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

	protected.POST("/budget", s.budgetHandler.CreateBudget)
	protected.GET("/budget", s.budgetHandler.GetAllBudgets)
	protected.GET("/budget/:budget_id", s.budgetHandler.GetBudgetByID)
	protected.PUT("/budget/:budget_id", s.budgetHandler.EditBudget)
	protected.DELETE("/budget/:budget_id", s.budgetHandler.DeleteBudget)
	protected.POST("/budget/:budget_id/share", s.budgetHandler.ShareBudget)

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
