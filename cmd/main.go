package main

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/adapters/handlers/httpv1"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/services"
	"github.com/zeabix-cloud-native/workshop-inventory-service/utils"

	"fmt"
	"log"
)

func main() {
	port := utils.GetEnv("PORT", "3000")
	initFile := utils.GetEnv("INIT_DB_FILE", "sample.json")

	repo := repository.NewMapDBRepository()
	log.Printf("InitDB from %s\n", initFile)
	utils.InitDB(initFile, repo)

	srv := services.NewInventoryService(repo)
	handler := httpv1.NewHTTPHandler(srv)

	log.Printf("Start inventory service at port: %s\n", port)
	handler.Serve(fmt.Sprintf(":%s", port))
}
