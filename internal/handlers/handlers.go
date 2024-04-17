package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mezink-records-server/internal/model/dto"
	"github.com/mezink-records-server/internal/service"
	"github.com/mezink-records-server/shared/response"
)

type Handler struct {
	Service service.Service
}

func ProvideHandler(service service.Service) Handler {
	return Handler{
		Service: service,
	}
}

func (h *Handler) Router(r chi.Router) {

	r.Route("/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/records", h.ViewRecords)
		})
	})
}

func (h *Handler) ViewRecords(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.Service.GetRecordList(request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}
