package fukui_bear_information

import "context"

type IFukuiBearInformation interface {
	Fetch(ctx context.Context) (*FukuiBearInformationLocation, error)
}
