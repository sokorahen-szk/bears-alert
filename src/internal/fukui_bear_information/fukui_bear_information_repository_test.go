package fukui_bear_information_test

import (
	"context"
	"testing"

	"github.com/sokorahen-szk/bears-alert/internal/fukui_bear_information"
)

func TestFukuiBearInformationRepositoryFetch(t *testing.T) {
	t.Skip("skip")
	fbir := fukui_bear_information.NewFukuiBearInformationRepository()
	fbir.Fetch(context.Background())
}
