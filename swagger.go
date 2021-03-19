package fiber_swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"path"
)

type Middleware struct {
	FilePath string
	BasePath string
}

func (s *Middleware) swaggerUIHandler(opts middleware.SwaggerUIOpts) http.Handler {
	return middleware.SwaggerUI(opts, nil)
}

func (s *Middleware) swaggerSpecFileHandler(swaggerUiHandler http.Handler) (http.Handler, error) {
	if _, err := os.Stat(s.FilePath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("%s file is not exist", s.FilePath))
	}

	specDoc, err := loads.Spec(s.FilePath)
	if err != nil {
		return nil, err
	}

	b, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
	if err != nil {
		return nil, err
	}

	return handlers.CORS()(middleware.Spec(s.BasePath, b, swaggerUiHandler)), nil
}

func (s *Middleware) Register(app *fiber.App) {
	swaggerUIOpts := middleware.SwaggerUIOpts{
		BasePath: s.BasePath,
		SpecURL:  path.Join(s.BasePath, "swagger.json"),
		Path:     "docs",
	}

	swaggerUiHandler := s.swaggerUIHandler(swaggerUIOpts)
	specFileHandler, err := s.swaggerSpecFileHandler(swaggerUiHandler)

	if err != nil {
		panic(err)
	}

	app.Use(path.Join(s.BasePath, swaggerUIOpts.Path), adaptor.HTTPHandler(swaggerUiHandler))
	app.Use(path.Join(s.BasePath, "swagger.json"), adaptor.HTTPHandler(specFileHandler))
}

func NewMiddleware(fileName string, basePath string) *Middleware {
	return &Middleware{
		FilePath: fileName,
		BasePath: basePath,
	}
}
