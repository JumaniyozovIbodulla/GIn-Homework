package main

import (
	"backend_course/lms/api"
	"backend_course/lms/config"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/service"
	"backend_course/lms/storage/postgres"
	"context"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		log.Error("error while connecting db, err: ", logger.Error(err))
		return
	}
	defer store.CloseDB()

	service := service.New(store, log)

	c := api.New(store, service, log)

	c.Run(":8080")
}
