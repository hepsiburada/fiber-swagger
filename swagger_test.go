package fiber_swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(method, target string, app *fiber.App) *http.Response {
	r := httptest.NewRequest(method, target, nil)

	resp, _ := app.Test(r, -1)

	return resp
}

func TestMiddleware_Register(t *testing.T) {
	t.Run("Endpoint check", func(t *testing.T) {
		app := fiber.New()

		middleware := NewMiddleware("./docs/swagger_test.json", "/")

		middleware.Register(app)

		w1 := performRequest("GET", "/docs", app)
		assert.Equal(t, 200, w1.StatusCode)

		w2 := performRequest("GET", "/swagger.json", app)
		assert.Equal(t, 200, w2.StatusCode)

		w3 := performRequest("GET", "/notfound", app)
		assert.Equal(t, 404, w3.StatusCode)
	})

	t.Run("Swagger.json file is not exist", func(t *testing.T) {
		app := fiber.New()

		middleware := NewMiddleware("./docs/swagger.json", "/")

		assert.Panics(t, func() {
			middleware.Register(app)
		}, "/swagger.json file is not exist")
	})

	t.Run("Swagger.json missing file", func(t *testing.T) {
		app := fiber.New()

		middleware := NewMiddleware("./docs/swagger_missing_test.json", "/")

		assert.Panics(t, func() {
			middleware.Register(app)
		}, "invalid character ':' after object key:value pair")
	})
}