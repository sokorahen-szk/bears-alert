package fukui_bear_information

import "context"

type fukuiBearInformationUsecase struct{}

func NewFukuiBearInformationUsecase() fukuiBearInformationUsecase {
	return fukuiBearInformationUsecase{}
}

func (fbiu fukuiBearInformationUsecase) Fetch(ctx context.Context) (*FukuiBearInformationLocation, error) {
	return nil, nil
}
