package data

import (
	"net/url"

	"github.com/LucasAMoen/balance-the-gathering/application"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var Cards = []application.Card{
	{
		Id:       uuid.MustParse("urn:uuid:0000419b-0bba-4488-8f7a-6194544ce91e"),
		Name:     "Forest",
		Url:      "https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e",
		Layout:   "normal",
		ImageUrl: url.URL{Path: "https://cards.scryfall.io/png/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.png?1721427487"},
		Price:    decimal.NewFromFloat32(0.32),
	},
	{
		Id:       uuid.MustParse("urn:uuid:e4a2d2c6-8eaa-4760-b620-921b807baa2e"),
		Name:     "Feather, the Redeemed",
		Url:      "https://api.scryfall.com/cards/e4a2d2c6-8eaa-4760-b620-921b807baa2e",
		Layout:   "normal",
		ImageUrl: url.URL{Path: "https://cards.scryfall.io/png/front/e/4/e4a2d2c6-8eaa-4760-b620-921b807baa2e.png?1557577142"},
		Price:    decimal.NewFromFloat32(0.71),
	},
}
