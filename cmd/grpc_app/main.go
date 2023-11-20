package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"task/pkg/grpc_handler/api"
	"task/pkg/repository"
	"task/pkg/serverity"
	"task/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializing config: %s", err.Error())
		return
	}
	logrus.Infof("config initialized successfully")

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env varibles: %s", err.Error())
		return
	}
	logrus.Infof("enviroment load successfully")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Fatal to connect to DB, because: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := api.NewGRPCHandlerTask(service)

	srv := new(serverity.GRPCServer)

	go func() {
		if err := srv.Run(handler); err != nil {
			logrus.Fatalf("Problem with start server, because %s", err.Error())
			return
		}
	}()
	logrus.Println("backend started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("backend shutting down")

	srv.ShutDown()

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
