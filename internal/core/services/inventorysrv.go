package services

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/ports"
)

type service struct {
	r ports.InventoryRepository
}

func NewInventoryService(r ports.InventoryRepository) ports.InventoryService {
	return &service{
		r: r,
	}
}

func (s *service) GetInventory(id string) (*domain.ItemInventory, error) {
	return s.r.GetInventoryByID(id)
}
