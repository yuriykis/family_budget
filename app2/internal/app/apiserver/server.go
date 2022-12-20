package apiserver

import (
	"app/internal/app/controllers"
	"app/internal/app/controllers/controller"
	"app/internal/app/middlewares"
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
	budgetController  controllers.IBudgetController
}

func newServer(store store.IStore) *server {
	s := &server{
		router:            gin.Default(),
		logger:            log.New(),
		userController:    controller.NewsUserController(store),
		serviceController: controller.NewServiceController(),
		budgetController:  controller.NewBudgetController(store),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {

	public := s.router.Group("/api")
	public.GET("/healthcheck", s.serviceController.CheckHealth)
	public.POST("/register", s.userController.RegisterUser)
	public.POST("/login", s.userController.AthenticateUser)

	protected := s.router.Group("/api/auth")
	protected.Use(middlewares.JwtMiddleware())
	protected.GET("/user", s.userController.GetUser)
	protected.POST("/budget", s.budgetController.CreateBudget)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	s.router.ServeHTTP(w, r)
}
