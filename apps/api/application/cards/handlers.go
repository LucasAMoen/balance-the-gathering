package cards

import (
	"log"
	"net/http"

	"github.com/LucasAMoen/balance-the-gathering/application/cards/data"
	"github.com/LucasAMoen/balance-the-gathering/application/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetCards(writer http.ResponseWriter, request *http.Request) {
	err := h.service.GetCards(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	json.Write(writer, http.StatusOK, data.Cards)
}

func (h *handler) GetCard(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	for idx := range data.Cards {
		if data.Cards[idx].Id.String() == id {
			json.Write(writer, http.StatusOK, data.Cards[idx])
			return
		}
	}
	json.Write(writer, http.StatusNotFound, nil)
}

func (h *handler) GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Healthy"))
}
