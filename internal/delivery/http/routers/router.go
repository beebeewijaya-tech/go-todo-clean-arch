package routers

import (
	"github.com/beebeewijaya-tech/go-todo/internal/delivery/http/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	router         *echo.Echo
	userController controllers.UserControllerInterface
	config         *viper.Viper
	privateRoute   *echo.Group
	todoContoller  controllers.TodoControllerInterface
}

func NewServer(
	config *viper.Viper,
	userController controllers.UserControllerInterface,
	todoContoller controllers.TodoControllerInterface,
) *Server {
	s := &Server{
		userController: userController,
		config:         config,
		todoContoller:  todoContoller,
	}

	s.router = echo.New()
	s.privateRoute = s.router.Group("")
	s.initMiddleware()
	s.initRoute()

	return s
}

func (s *Server) initMiddleware() {
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	s.router.Use(middleware.Recover())
	s.authMiddleware()
}

func (s *Server) authMiddleware() {
	s.privateRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.config.GetString("jwt.secret")),
	}))
}

func (s *Server) initRoute() {
	s.initAuthRoute()
	s.initTodoRoutes()
}

func (s *Server) initAuthRoute() {
	s.router.POST("/api/login", s.userController.Login)
	s.router.POST("/api/register", s.userController.Register)
}

func (s *Server) initTodoRoutes() {
	s.privateRoute.GET("/api/todo/list", s.todoContoller.ListTodo)
	s.privateRoute.GET("/api/todo/:id", s.todoContoller.GetTodo)
	s.privateRoute.POST("/api/todo", s.todoContoller.CreateTodo)
	s.privateRoute.PUT("/api/todo/:id", s.todoContoller.UpdateTodo)
	s.privateRoute.DELETE("/api/todo/:id", s.todoContoller.DeleteTodo)
}

func (s *Server) StartServer(addr string) {
	s.router.Logger.Fatal(s.router.Start(addr))
}
