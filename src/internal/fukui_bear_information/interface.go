package fukui_bear_information

import "context"

type IFukuiBearInformationRepository interface {
	Fetch(ctx context.Context) ([]*FukuiBearInformation, error)
}
