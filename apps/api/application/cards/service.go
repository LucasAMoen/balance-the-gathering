package cards

import (
	"context"
	"log"
	"net/url"

	"github.com/LucasAMoen/balance-the-gathering/application"
	repository "github.com/LucasAMoen/balance-the-gathering/infrastructure/adapters/postgresql/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type svc struct {
	repository repository.Querier
}

func NewService(repository repository.Querier) Service {
	return &svc{
		repository: repository,
	}
}

type Service interface {
	GetCards(ctx context.Context) ([]application.Card, error)
	GetCardById(ctx context.Context, id uuid.UUID) (application.Card, error)
}

func (s *svc) GetCards(ctx context.Context) ([]application.Card, error) {
	repositoryCards, error := s.repository.GetCards(ctx)
	var cardArray = []application.Card{}
	for _, card := range repositoryCards {
		cardArray = append(cardArray, toCard(card))
	}
	return cardArray, error
}

func (s *svc) GetCardById(ctx context.Context, id uuid.UUID) (application.Card, error) {
	repositoryCard, error := s.repository.GetCardById(ctx, pgtype.UUID{Bytes: id, Valid: true})
	log.Println("Here")
	return toCard(repositoryCard), error
}

func toCard(card repository.Card) application.Card {
	return application.Card{
		Id:        uuid.MustParse("urn:uuid:" + card.ID.String()),
		Name:      card.Name,
		Url:       url.URL{Path: card.Url},
		Layout:    card.Layout.String,
		ImageUrl:  url.URL{Path: card.Imageurl.String},
		Price:     decimal.NewFromFloat32(float32(card.Price.Exp)),
		CreatedAt: card.Createdat.Time,
		UpdatedAt: card.Updatedat.Time,
	}
}
