package repositories

import (
	"context"

	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/domain"
)

type bearsDetailRepository struct {
}

func NewBearsDetailRepository() bearsDetailRepository {
	return bearsDetailRepository{}
}

func (bdr bearsDetailRepository) Insert(context.Context, *domain.BearsDetail) error {
	return nil
}

func (bdr bearsDetailRepository) List(context.Context) ([]domain.BearsDetail, error) {
	return nil, nil
}
