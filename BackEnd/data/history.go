package data

// TODO: history

type HistoryRespBody struct {
	RespBody
	Data HistoryRespBodyIn `json:"data"`
}

type HistoryRespBodyIn struct {
	Cursor Hcursor  `json:"cursor"`
	List   []ListOb `json:"list"`
}

type Hcursor struct {
	Max      int    `json:"max"`
	View_at  int    `json:"view_at"`
	Business string `json:"business"`
	Ps       int    `json:"ps"`
}

type ListOb struct {
	History History `json:"history"`
	View_at int     `json:"view_at"`
}

type History struct {
	Bvid string `json:"bvid"`
}

type LastViewAt struct {
	Id    int `gorm:"primaryKey"`
	Value int
}
