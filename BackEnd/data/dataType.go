package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Vgroup interface {
	GetVideo() ([]Video, error)
	GetLabel() Label
	SetLabel(Label)
	GetStr() (string, error)
}

type VgType int

const (
	IsSeason VgType = iota
	IsUp
)

// up
type Up struct {
	UpID  int    `json:"mid"`
	Name  string `json:"uname"`
	Label Label  `json:"label"`
	Type  VgType `json:"type"`
}

func (u Up) GetVideo() ([]Video, error) {
	// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E6%8A%95%E7%A8%BF%E8%A7%86%E9%A2%91%E6%98%8E%E7%BB%86
	url := "https://api.bilibili.com/x/space/wbi/arc/search?mid=" + strconv.Itoa(u.UpID)
	var jsonData VideoRespBody
	err := Get(url, &jsonData)
	// TODO: 优化，不要每次都从网上中获取视频
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}

	listp := &jsonData.Data.List.Vlist
	for i := range *listp { // 注意 for range 会复制元素
		(*listp)[i].Label = u.Label
		(*listp)[i].UpId = u.UpID
		(*listp)[i].UpName = u.Name
	}
	return *listp, nil
}

func (u *Up) GetLabel() Label {
	return u.Label
}

func (u *Up) SetLabel(l Label) {
	u.Label = l
}

func (u *Up) GetStr() (string, error) {
	b, e := json.Marshal(*u)
	return string(b), e
}

// season
type Season struct {
	UpID     int    `json:"mid"`
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

// label
type Label int

const (
	NoLab      Label = iota
	Ignore           // 1
	Normal           // 2
	Prefer           // 3
	VeryPrefer       // 4
)

// video
type Video struct {
	Bvid      string `json:"bvid"`
	Title     string `json:"title"`
	View      int    `json:"play"`
	LengthStr string `json:"length"` // don't use => Length()
	LengthS   int    // don't use => Length()
	Pic       string `json:"pic"`
	Label     Label
	Seen      bool
	UpId      int    `json:"mid"`
	UpName    string `json:"up_name"`
}

func (v *Video) Length() int { // return length in second
	if v.LengthS == 0 {
		t := strings.Split(v.LengthStr, ":")
		if len(t) != 2 {
			return -1
		}
		i1, err := strconv.Atoi(t[0])
		i2, err2 := strconv.Atoi(t[1])
		if err != nil || err2 != nil {
			return -1
		}
		v.LengthS = 60*i1 + i2
	}
	return v.LengthS
}
