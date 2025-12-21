package data

import (
	"encoding/json"
	"fmt"
	"strconv"
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

func (u Up) GetVideo() ([]Video, error) {
	// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html#%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7%E6%8A%95%E7%A8%BF%E8%A7%86%E9%A2%91%E6%98%8E%E7%BB%86
	url := "https://api.bilibili.com/x/space/wbi/arc/search?mid=" + strconv.Itoa(u.UpID)
	var jsonData VideoRespBody
	err := Get(url, &jsonData)
	// TODO: 优化，不要每次都从网上中获取视频
	if err != nil {
		return nil, fmt.Errorf("error get up video: %s", err)
	}

	listp := &jsonData.Data.List.Vlist
	for i := range *listp { // 注意 for range 会复制元素
		(*listp)[i].Label = u.Label
		(*listp)[i].UpId = u.UpID
		(*listp)[i].UpName = u.Name
	}
	return *listp, nil
}
