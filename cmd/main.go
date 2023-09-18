package main

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/adapters/handlers/httpv1"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/services"
	"github.com/zeabix-cloud-native/workshop-inventory-service/utils"
)

func main() {
	repo := repository.NewMapDBRepository()
	utils.InitDB("sample.json", repo)

	srv := services.NewInventoryService(repo)
	handler := httpv1.NewHTTPHandler(srv)
	handler.Serve(":3030")
}
