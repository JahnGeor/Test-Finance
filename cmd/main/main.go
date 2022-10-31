package main

import (
	"context"
	"github.com/jahngeor/avito-tech/internal/controller"
	"github.com/jahngeor/avito-tech/internal/gateway"
	"github.com/jahngeor/avito-tech/internal/service"
	"github.com/jahngeor/avito-tech/pkg/config"
	"github.com/jahngeor/avito-tech/pkg/server"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	configs, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Ошибка во время чтения конфигурации: %s", err.Error())
	}

	db, err := gateway.NewPostgresDB(context.Background(), configs)
	if err != nil {
		logrus.Fatalf("Ошибка при подключении к БД: %s", err.Error())
	}

	if err := gateway.TestPing(db, context.Background()); err != nil {
		logrus.Fatalf("Ошибка во время пинга БД: %s", err.Error())
	}
	gtw := gateway.NewGateway(db)
	services := service.NewServices(gtw)
	ctrl := controller.NewController(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(configs.ServerConfig.Port, ctrl.InitRoutes()); err != nil {
			logrus.Fatalf("Ошибка во время запуска http-сервера: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Приложение запущено.")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Ошибка в завершении работы сервера: %s", err.Error())
	}

	db.Close()
}
