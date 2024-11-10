package data

import (
	"strconv"
	"strings"
)

// video
type Video struct {
	Bvid      string `json:"bvid" gorm:"primaryKey;type:char(12)"` // 主键
	Title     string `json:"title"`
	View      int    `json:"play"`
	LengthStr string `json:"length"` // don't use => Length()
	LengthS   int    // don't use => Length()
	Pic       string `json:"pic"`
	Label     Label
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

type SeenVideo struct {
	Bvid string `json:"bvid" gorm:"primaryKey;type:char(12)"` // 主键
}
