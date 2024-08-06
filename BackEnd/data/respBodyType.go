package data

// respBody
type RespBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TagRespBody struct {
	RespBody
	Data []Uptag `json:"data"`
}

// up
type UpRespBody struct {
	RespBody
	Data []Up `json:"data"`
}

type VideoRespBody struct {
	RespBody
	Data VideoRespBodyIn `json:"data"`
}

type VideoRespBodyIn struct {
	List VideoRespBodyIn2 `json:"list"`
}

type VideoRespBodyIn2 struct {
	Vlist []Video `json:"vlist"`
}

type Uptag struct {
	Tagid int    `json:"tagid"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type DetailEachTag struct {
	TagInfo Uptag `json:"taginfo"`
	UpList  []Up  `json:"uplist"`
}

// season
type SeasonRespBody struct {
	RespBody
	Data SeasonRespBodyIn `json:"data"`
}

type SeasonRespBodyIn struct {
	List []SeasonVideo `json:"archives"`
	Meta SeasonDetail  `json:"meta"`
}

type SeasonVideo struct {
	Bvid    string `json:"bvid"`
	Title   string `json:"title"`
	Stat    Svstat `json:"stat"`
	LengthS int    `json:"duration"`
	Pic     string `json:"pic"`
}

type Svstat struct {
	View int `json:"view"`
}

type SeasonDetail struct {
	SeasonID int    `json:"season_id"`
	Name     string `json:"name"`
	UpID     int    `json:"mid"`
}

// SubscriptSeason
type SubSeasonRespBody struct {
	RespBody
	Data SubSeasonRespBodyIn `json:"data"`
}

type SubSeasonRespBodyIn struct {
	List []Season `json:"list"`
}
