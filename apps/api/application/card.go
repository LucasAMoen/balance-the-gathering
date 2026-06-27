package application

import (
	"net/url"
	"time"

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
