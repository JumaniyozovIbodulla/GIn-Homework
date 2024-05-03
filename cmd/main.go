package main

import (
	"backend_course/lms/api"
	"backend_course/lms/config"
	"backend_course/lms/storage/postgres"
	"context"
	"fmt"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	c := api.New(store)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8080")
}
