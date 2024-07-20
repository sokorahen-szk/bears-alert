package usecases

import (
	"context"

	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/domain"
	"github.com/sokorahen-szk/bears-alert/internal/fukui_bear_information"
)

type GetBearsDetailUsecase struct {
	bearsDetailRepo                domain.IBearsDetailRepository
	fukuiBearInformationRepository fukui_bear_information.IFukuiBearInformationRepository
}

func NewGetBearsDetailUsecase(
	bearsDetailRepo domain.IBearsDetailRepository,
	fukuiBearInformationRepository fukui_bear_information.IFukuiBearInformationRepository,
) *GetBearsDetailUsecase {
	return &GetBearsDetailUsecase{
		bearsDetailRepo:                bearsDetailRepo,
		fukuiBearInformationRepository: fukuiBearInformationRepository,
	}
}

func (gbdu GetBearsDetailUsecase) Exec(ctx context.Context) error {
	return nil
}
