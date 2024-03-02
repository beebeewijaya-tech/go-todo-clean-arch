package main

import (
	"github.com/beebeewijaya-tech/go-todo/internal/config"
	"github.com/beebeewijaya-tech/go-todo/internal/db"
	"github.com/beebeewijaya-tech/go-todo/internal/delivery/http/controllers"
	"github.com/beebeewijaya-tech/go-todo/internal/delivery/http/routers"
	"github.com/beebeewijaya-tech/go-todo/internal/repositories/postgres"
	"github.com/beebeewijaya-tech/go-todo/internal/usecases"
	"github.com/beebeewijaya-tech/go-todo/internal/utils"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogger(viper)
	database := db.NewDatabase(viper, log)
	err := database.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// Utils
	utilities := utils.NewUtils(viper)

	// Repository
	userRepo := postgres.NewUserRepository(database)
	todoRepo := postgres.NewTodoRepository(database)

	// Usecase
	userUsecase := usecases.NewUserUsecase(userRepo, utilities)
	todoUsecase := usecases.NewTodoUsecase(todoRepo)

	// Controllers
	userController := controllers.NewUserController(userUsecase)
	todoContoller := controllers.NewTodoController(todoUsecase, utilities)

	// Router
	router := routers.NewServer(viper, userController, todoContoller)
	router.StartServer(":9000")
}
