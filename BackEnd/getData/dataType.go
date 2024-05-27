package getData

type respBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type tagRespBody struct {
	respBody
	Data []uptag `json:"data"`
}

type upRespBody struct {
	respBody
	Data []up `json:"data"`
}

type uptag struct {
	Tagid int    `json:"tagid"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type up struct {
	UpID int    `json:"mid"`
	Name string `json:"uname"`
}

type detailEachTag struct {
	TagInfo uptag `json:"taginfo"`
	UpList  []up  `json:"uplist"`
}
