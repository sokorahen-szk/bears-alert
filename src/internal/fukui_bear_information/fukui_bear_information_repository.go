package fukui_bear_information

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const bearInformationURL = "https://tsukinowaguma.pref.fukui.lg.jp"

type FukuiBearInformationRepository struct{}

func NewFukuiBearInformationRepository() IFukuiBearInformationRepository {
	return new(FukuiBearInformationRepository)
}

func (fbir FukuiBearInformationRepository) Fetch(ctx context.Context) ([]*FukuiBearInformation, error) {
	res, err := http.Get(bearInformationURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	fukuiBearInformations := make([]*FukuiBearInformation, 0)

	doc.Find(".container #HeaderPlace_gdvList").Each(func(i int, row *goquery.Selection) {
		row.Find("tr").Each(func(_ int, tr *goquery.Selection) {
			tdList := make([]string, 0)
			tr.Find("td").Each(func(i int, td *goquery.Selection) {
				tdList = append(tdList, td.Text())
			})
			if len(tdList) == 7 {
				fukuiBearInformations = append(fukuiBearInformations, &FukuiBearInformation{
					ID:   tdList[0],
					Kind: tdList[1],
					Location: FukuiBearInformationLocation{
						Prefectures:    DefaultPrefectures,
						Municipalities: tdList[2],
						Detail:         tdList[3],
					},
					Datetime:                 fmt.Sprintf("%s %s", tdList[4], tdList[5]),
					EyewitnessClassification: tdList[6],
				})
			}
		})
	})

	return fukuiBearInformations, nil
}
