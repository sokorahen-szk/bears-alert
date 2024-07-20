package fukui_bear_information

type FukuiBearInformation struct {
	Kind                     string
	Location                 FukuiBearInformationLocation
	EyewitnessClassification string
	Datetime                 string
}

type FukuiBearInformationLocation struct {
	Prefectures    string
	Municipalities string
}
