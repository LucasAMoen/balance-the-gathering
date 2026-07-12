package application

import (
	"context"

	"github.com/LucasAMoen/balance-the-gathering/domain"
	"github.com/google/uuid"
)

type IRepository interface {
	GetCards(ctx context.Context) ([]domain.Card, error)
	GetCardById(ctx context.Context, id uuid.UUID) (domain.Card, error)
}
