package repositories

import (
	"context"

	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/domain"
)

type BearsDetailRepository struct {
}

func NewBearsDetailRepository() domain.IBearsDetailRepository {
	return new(BearsDetailRepository)
}

func (bdr BearsDetailRepository) Insert(context.Context, *domain.BearsDetail) error {
	return nil
}

func (bdr BearsDetailRepository) List(context.Context) ([]*domain.BearsDetail, error) {
	return nil, nil
}
