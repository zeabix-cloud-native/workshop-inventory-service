package httpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/zeabix-cloud-native/workshop-inventory-service/internal/core/ports"

	"encoding/json"
)

type InventoryResponse struct {
	ID     string `json:"id"`
	Amount uint   `json:"amount"`
}

type Handler struct {
	s   ports.InventoryService
	app *fiber.App
}

func NewHTTPHandler(srv ports.InventoryService) *Handler {
	return &Handler{
		s:   srv,
		app: fiber.New(),
	}
}

func (h *Handler) Serve(port string) error {
	h.app.Use(cors.New())
	g := h.app.Group("inventory-service")
	v1 := g.Group("v1")
	v1.Get("/inventories/:id", h.getInventoryHandler)

	return h.app.Listen(port)
}

func (h *Handler) getInventoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}

	i, err := h.s.GetInventory(id)

	if err == ports.ErrItemNotFound {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	res := new(InventoryResponse)
	res.ID = i.ID
	res.Amount = i.Amount

	resStr, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK).SendString(string(resStr))
	return nil
}
