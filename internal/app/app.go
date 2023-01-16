package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"go.service/internal/app/account"

	"go.service/internal/pkg/server"
	"go.service/internal/pkg/util"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func Run() {
	config := new(Config)
	if err := util.LoadConfig(".", "app", config); err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := util.ConnectDb(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	e := echo.New()

	accountRepository := account.RepositoryInterface(account.NewRepository(conn))
	accountUseCase := account.UseCaseInterface(account.NewUseCase(&accountRepository))
	accountController := account.ControllerInterface(account.NewController(&accountUseCase))
	account.Handle(e, &accountController)

	// healthcheck.Handle(e)

	if err := server.New(e).Run(); err != nil {
		log.Fatal("server error:", err)
	}
}
