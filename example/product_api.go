package main

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"sync"
)

var (
	productHandlerOnce sync.Once
	productHandler     *handler
)

type handler struct {
}

// swagger:parameters getProduct
type productRequest struct {
	// Id of an product
	// In: path
	ProductId string `json:"productId"`
}

// swagger:route GET /product/{productId} Product getProduct
// responses:
//  200: ProductModel
func (h *handler) GetByProductId(c *fiber.Ctx) error {
	productId := c.Params("productId")
	productId, err := url.QueryUnescape(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product := NewProductModel("Test Product")

	_ = c.JSON(product)

	c.Status(fiber.StatusOK)
	return nil
}

func SetupRoutes(app *fiber.App) {
	app.Get("/product/:productId", productHandler.GetByProductId)
}


func GetHandler() *handler {
	productHandlerOnce.Do(func() {
		productHandler = NewHandler()
	})

	return productHandler
}

func NewHandler() *handler {
	return &handler{}
}
