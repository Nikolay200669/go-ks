package http

import (
	"net/http"

	"github.com/Nikolay200669/go-ks/internal/app/application"
	"github.com/julienschmidt/httprouter"
)

// SetupRouter configures the router
func SetupRouter(service *application.CalculateService) http.Handler {
	router := httprouter.New()
	router.POST("/calculate", ValidateInput(CalculateHandler(service)))
	return router
}
