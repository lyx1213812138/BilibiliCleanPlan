package data

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Season struct {
	UpID     int    `json:"mid" gorm:"primaryKey;"` // 主键
	SeasonID int    `json:"id"`
	Name     string `json:"title"`
	Label    Label
	Type     VgType `json:"type"`
}

func (s *Season) GetVideo() ([]Video, error) {
	// https://socialsisteryi.github.io/bilibili-API-collect/docs/video/collection.html
	// https://api.bilibili.com/x/polymer/web-space/seasons_archives_list?mid=1567748478&season_id=32744
	url := "https://api.bilibili.com/x/polymer/web-space/seasons_archives_list?mid=" + strconv.Itoa(s.UpID) + "&season_id=" + strconv.Itoa(s.SeasonID)
	var jsonData SeasonRespBody
	err := Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}

	s.Name = jsonData.Data.Meta.Name
	var video []Video
	for _, sv := range jsonData.Data.List {
		video = append(video, Video{
			Bvid:    sv.Bvid,
			Title:   sv.Title,
			View:    sv.Stat.View,
			Pic:     sv.Pic,
			LengthS: sv.LengthS,
			Label:   s.Label,
			UpId:    s.UpID,
			UpName:  s.Name,
		})
	}

	return video, nil
}

func (s *Season) GetLabel() Label {
	return s.Label
}

func (s *Season) SetLabel(l Label) {
	s.Label = l
}

func (u *Season) GetStr() (string, error) {
	b, e := json.Marshal(*u)
	return string(b), e
}

func (u *Season) GetUpID() int {
	return u.UpID
}
