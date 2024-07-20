package domain

type BearsDetail struct {
	Kind                     string
	Location                 BearsDetailLocation
	EyewitnessClassification string
	Datetime                 string
	IsNotice                 bool
	RemindNoticeCount        int8
	MaxNoticeCount           int8
}

type BearsDetailLocation struct {
	Prefectures    string
	Municipalities string
}
