package myType

type RespBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TagRespBody struct {
	RespBody
	Data []Uptag `json:"data"`
}

type UpRespBody struct {
	RespBody
	Data []Up `json:"data"`
}

type Uptag struct {
	Tagid int    `json:"tagid"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Up struct {
	UpID  int    `json:"mid"`
	Name  string `json:"uname"`
	Label Label  `json:"label"`
}

type DetailEachTag struct {
	TagInfo Uptag `json:"taginfo"`
	UpList  []Up  `json:"uplist"`
}

type Label int

const (
	NoLab      Label = iota
	Ignore           // 1
	Normal           // 2
	Prefer           // 3
	VeryPrefer       // 4
)
