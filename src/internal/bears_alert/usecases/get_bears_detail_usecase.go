package usecases

import (
	"context"

	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/domain"
	"github.com/sokorahen-szk/bears-alert/internal/fukui_bear_information"
)

type getBearsDetailUsecase struct {
	bearsDetailRepo      domain.IBearsDetailRepository
	fukuiBearInformation fukui_bear_information.IFukuiBearInformation
}

func NewGetBearsDetailUsecase(
	bearsDetailRepo domain.IBearsDetailRepository,
	fukuiBearInformation fukui_bear_information.IFukuiBearInformation,
) getBearsDetailUsecase {
	return getBearsDetailUsecase{
		bearsDetailRepo:      bearsDetailRepo,
		fukuiBearInformation: fukuiBearInformation,
	}
}

func (gbdu getBearsDetailUsecase) Exec(ctx context.Context) error {
	return nil
}
