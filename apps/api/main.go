package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Card struct {
	Id        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Url       string          `json:"url"`
	Layout    string          `json:"layout"`
	ImageUrl  url.URL         `json:"imageUrl"`
	Allowed   bool            `json:"allowed"`
	Price     decimal.Decimal `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
}

var cards = []Card{
	{
		Id:       uuid.MustParse("urn:uuid:0000419b-0bba-4488-8f7a-6194544ce91e"),
		Name:     "Forest",
		Url:      "https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e",
		Layout:   "normal",
		ImageUrl: url.URL{Path: "https://cards.scryfall.io/png/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.png?1721427487"},
		Allowed:  true,
		Price:    decimal.NewFromFloat32(0.32),
	},
	{
		Id:       uuid.MustParse("urn:uuid:e4a2d2c6-8eaa-4760-b620-921b807baa2e"),
		Name:     "Feather, the Redeemed",
		Url:      "https://api.scryfall.com/cards/e4a2d2c6-8eaa-4760-b620-921b807baa2e",
		Layout:   "normal",
		ImageUrl: url.URL{Path: "https://cards.scryfall.io/png/front/e/4/e4a2d2c6-8eaa-4760-b620-921b807baa2e.png?1557577142"},
		Allowed:  true,
		Price:    decimal.NewFromFloat32(0.71),
	},
}

func getCards(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cards)
}

func main() {
	router := gin.Default()
	router.GET("/cards", getCards)
	router.Run("localhost:8080")
}
