package apiserver

import (
	"app/internal/app/graph"
	"app/internal/app/handlers"
	"app/internal/app/handlers/handler"
	"app/internal/app/middlewares"
	"app/internal/app/store"

	"net/http"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router             *gin.Engine
	logger             *log.Logger
	graphqlEndpoint    *gqlhandler.Server
	userHandler        handlers.IUserHandler
	serviceHandler     handlers.IServiceHandler
	budgetHandler      handlers.IBudgetHandler
	categoryHandler    handlers.ICategoryHandler
	transactionHandler handlers.ITransactionHandler
}

func newServer(store store.IStore) *server {
	s := &server{
		router:             gin.Default(),
		logger:             log.New(),
		graphqlEndpoint:    gqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(store)})),
		userHandler:        handler.NewsUserHandler(store),
		serviceHandler:     handler.NewServiceHandler(),
		budgetHandler:      handler.NewBudgetHandler(store),
		categoryHandler:    handler.NewCategoryHandler(store),
		transactionHandler: handler.NewTransactionHandler(store),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	s.router.Use(middlewares.CorsMiddleware())

	public := s.router.Group("/api")
	public.GET("/healthcheck", s.serviceHandler.CheckHealth)
	public.POST("/register", s.userHandler.RegisterUser)
	public.POST("/login", s.userHandler.AthenticateUser)

	protected := s.router.Group("/api/auth")
	protected.Use(middlewares.JwtMiddleware())
	protected.GET("/user", s.userHandler.GetUser)
	protected.GET("/user/:user_id", s.userHandler.GetUserByID)

	protected.POST("/budget", s.budgetHandler.CreateBudget)
	protected.GET("/budget", s.budgetHandler.GetAllBudgets)
	protected.GET("/budget/:budget_id", s.budgetHandler.GetBudgetByID)
	protected.PUT("/budget/:budget_id", s.budgetHandler.EditBudget)
	protected.DELETE("/budget/:budget_id", s.budgetHandler.DeleteBudget)
	protected.POST("/budget/:budget_id/transaction", s.transactionHandler.CreateTransaction)
	protected.GET("/budget/:budget_id/transaction", s.transactionHandler.GetTransactions)
	protected.POST("/budget/:budget_id/share", s.budgetHandler.ShareBudget)

	protected.POST("/category", s.categoryHandler.CreateCategory)
	protected.GET("/category", s.categoryHandler.FindAllCategories)

	// graphql endpoint
	s.router.POST("/query", func(c *gin.Context) {
		s.graphqlEndpoint.ServeHTTP(c.Writer, c.Request)
	})
	// graphql playground
	s.router.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
