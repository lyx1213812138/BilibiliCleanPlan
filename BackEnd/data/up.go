package data

import (
	"encoding/json"
)

type Up struct {
	UpID  int    `json:"mid" gorm:"primaryKey;"` // 主键
	Name  string `json:"uname"`
	Label Label  `json:"label"`
	Type  VgType `json:"type"`
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

func (u *Up) GetUpID() int {
	return u.UpID
}
