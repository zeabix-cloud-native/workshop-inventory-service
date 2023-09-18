package utils

import (
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/ports"

	"io/ioutil"
	"os"

	"encoding/json"
)

type Inventories struct {
	Inventories []domain.ItemInventory `json:"inventories"`
}

type Inventory struct {
	ID     string `json:"id"`
	Amount uint   `json:"amount"`
}

func InitDB(file string, repo ports.InventoryRepository) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var inventories Inventories
	json.Unmarshal(byteValue, &inventories)

	for _, v := range inventories.Inventories {
		inv := new(domain.ItemInventory)
		inv.ID = v.ID
		inv.Amount = v.Amount

		err := repo.Save(inv)
		if err != nil {
			return err
		}
	}

	return nil

}
