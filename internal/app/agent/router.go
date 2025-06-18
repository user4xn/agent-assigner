package agent

import (
	"agent-assigner/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func (h *handler) WebhookRouter(r chi.Router) {
	// add timing middleware for tracking time
	r.Use(middleware.Timing)
	r.Post("/webhook-assign", h.WebhookAssigment)
}
