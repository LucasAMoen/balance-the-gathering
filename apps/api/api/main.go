package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

func getCard(context *gin.Context) {
	id := context.Query("id")
	for idx := range cards {
		if cards[idx].Id.String() == id {
			context.IndentedJSON(http.StatusOK, cards[idx])
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, nil)
}

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost"
	}
	fullAddress := "http://" + serverAddress

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", fullAddress + ":5173"}
	config.AllowHeaders = []string{"Accept", "Access-Control-Allow-Origin", "Referer", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "user-agent"}

	router.Use(cors.New(config))
	router.GET("/cards", getCards)
	router.GET("/card", getCard)
	router.Run(serverAddress + ":8080")
}
