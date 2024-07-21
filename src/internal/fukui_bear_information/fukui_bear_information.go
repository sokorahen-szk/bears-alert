package fukui_bear_information

const DefaultPrefectures = "福井県"

type FukuiBearInformation struct {
	ID                       string
	Kind                     string
	Location                 FukuiBearInformationLocation
	Datetime                 string
	EyewitnessClassification string
}

type FukuiBearInformationLocation struct {
	Prefectures    string
	Municipalities string
	Detail         string
}
