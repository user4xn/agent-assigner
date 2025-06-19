package agent

import (
	"agent-assigner/internal/factory"
	"agent-assigner/pkg/util"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) WebhookAssigment(w http.ResponseWriter, r *http.Request) {
	// call webhook assignment function from service
	err := h.service.WebhookAssigment(r.Context())
	if err != nil {
		log.Println(err)
		response := util.APIResponse("error", 500, "an error occured: "+err.Error(), nil)
		util.JSON(w, http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("success", 200, "success run webhook assignment", nil)
	util.JSON(w, http.StatusOK, response)
}
