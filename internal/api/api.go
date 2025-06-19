package api

import (
	"agent-assigner/internal/app/agent"
	"agent-assigner/internal/dto"
	"agent-assigner/internal/factory"
	"agent-assigner/pkg/util"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// parent routing system
func NewAPI(r *chi.Mux, f *factory.Factory) {
	// default index route return service info
	r.Get("/", Index)

	// grouping v1 api passing to individual routes each app
	r.Route("/api/v1/agent", func(r chi.Router) {
		agent.NewHandler(f).WebhookRouter(r)
	})

	// handle not found routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		util.JSON(w, http.StatusNotFound, util.APIResponse("error", 404, "route not found", nil))
	})

	// handle invalid method on routes
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		util.JSON(w, http.StatusMethodNotAllowed, util.APIResponse("error", 405, "method not allowed", nil))
	})

	PrintRoutes(r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	info := dto.ServerInfo{
		ServiceName: "agent-assigner",
		Version:     "1.0.0",
	}

	util.JSON(w, http.StatusOK, info)
}

func PrintRoutes(r chi.Routes) {
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s] %s (middlewares: %d)\n", method, route, len(middlewares))
		return nil
	})
}
