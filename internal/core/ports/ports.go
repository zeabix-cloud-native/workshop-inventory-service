package ports

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/domain"

	"errors"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type InventoryService interface {
	GetInventory(id string) (*domain.ItemInventory, error)
}

type InventoryRepository interface {
	Save(i *domain.ItemInventory) error
	GetInventoryByID(id string) (*domain.ItemInventory, error)
}
