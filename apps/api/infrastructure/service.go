package infrastructure

import (
	"context"
	"net/url"

	"github.com/LucasAMoen/balance-the-gathering/application"
	"github.com/LucasAMoen/balance-the-gathering/domain"
	repository "github.com/LucasAMoen/balance-the-gathering/infrastructure/adapters/postgresql/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type svc struct {
	repository repository.Querier
}

func NewRepository(repository repository.Querier) application.IRepository {
	return &svc{
		repository: repository,
	}
}

func (s *svc) GetCards(ctx context.Context) ([]domain.Card, error) {
	repositoryCards, error := s.repository.GetCards(ctx)
	var cardArray = []domain.Card{}
	for _, card := range repositoryCards {
		cardArray = append(cardArray, toCard(card))
	}
	return cardArray, error
}

func (s *svc) GetCardById(ctx context.Context, id uuid.UUID) (domain.Card, error) {
	repositoryCard, error := s.repository.GetCardById(ctx, pgtype.UUID{Bytes: id, Valid: true})
	return toCard(repositoryCard), error
}

func toCard(card repository.Card) domain.Card {
	return domain.Card{
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
