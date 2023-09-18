package repository

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/ports"
)

type mapdb struct {
	DB map[string]uint
}

func NewMapDBRepository() ports.InventoryRepository {
	return &mapdb{
		DB: make(map[string]uint),
	}
}

func (db *mapdb) Save(i *domain.ItemInventory) error {
	db.DB[i.ID] = i.Amount
	return nil
}

func (db *mapdb) GetInventoryByID(id string) (*domain.ItemInventory, error) {
	i, ok := db.DB[id]
	if !ok {
		return nil, ports.ErrItemNotFound
	}

	item := new(domain.ItemInventory)
	item.ID = id
	item.Amount = i

	return item, nil
}
