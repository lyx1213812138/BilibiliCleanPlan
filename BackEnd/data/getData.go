package data

import (
	"fmt"
	"os"
	"strconv"
)

var myVmid string

func init() {
	vmid, err := os.ReadFile("./myVmid.txt")
	// 相对路径是相对于当前工作目录的。当前工作目录是指程序启动时所在的目录
	if err != nil {
		fmt.Printf("open vmid file failed: %s\n", err)
	}
	myVmid = string(vmid)
}

func GetSubscriptTags() ([]Uptag, error) {
	url := "https://api.bilibili.com/x/relation/tags?vmid=" + myVmid
	var jsonData TagRespBody
	err := Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func GetSubscriptUpByTag(tagid int) ([]Up, error) {
	url := "https://api.bilibili.com/x/relation/tag?tagid=" + strconv.Itoa(tagid)
	var jsonData UpRespBody
	err := Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func GetSubSeason() ([]Season, error) {
	url := "https://api.bilibili.com/x/v3/fav/folder/collected/list?pn=1&ps=20&up_mid=" + myVmid + "&platform=web"
	var jsonData SubSeasonRespBody
	err := Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get subseason from %s: %s", url, err)
	}
	var season []Season
	for _, s := range jsonData.Data.List {
		season = append(season, Season{
			UpID:     s.UpID,
			SeasonID: s.SeasonID,
			Name:     s.Name,
			Label:    NoLab,
		})
	}
	return season, nil
}
