package cards

import (
	"log"
	"net/http"

	"github.com/LucasAMoen/balance-the-gathering/application/json"
	"github.com/google/uuid"
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
	cards, err := h.service.GetCards(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	json.Write(writer, http.StatusOK, cards)
}

func (h *handler) GetCard(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	card, error := h.service.GetCardById(request.Context(), uuid.MustParse("urn:uuid:"+id))

	if error != nil {
		log.Println(error)
		http.Error(writer, error.Error(), http.StatusInternalServerError)
	}

	json.Write(writer, http.StatusOK, card)
}

func (h *handler) GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Healthy"))
}
