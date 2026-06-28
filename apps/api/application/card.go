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
	Url       url.URL         `json:"url"`
	Layout    string          `json:"layout"`
	ImageUrl  url.URL         `json:"imageUrl"`
	Price     decimal.Decimal `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
