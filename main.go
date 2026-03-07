package main

import (
	"fmt"
	"os"

	"github.com/GeaRvant/go_final_project/pkg/api"
	"github.com/GeaRvant/go_final_project/pkg/db"
	"github.com/GeaRvant/go_final_project/pkg/server"
)

func main() {
	webDir := "./web"
	port := ":7540"

	err := db.Init("scheduler.db")
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}

	api.Init()

	server.StartServer(port, webDir)
}
