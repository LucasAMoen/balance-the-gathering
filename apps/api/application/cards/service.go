package cards

import "context"

type Service interface {
	GetCards(ctx context.Context) error
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) GetCards(ctx context.Context) error {
	return nil
}
