// get data from bilibili api using get.go
package get

import (
	"fmt"
	"strconv"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
)

var myVmid string

// LEARN : init函数在main函数执行之前，自动被调用
func SetMyVmid(vmid float64) {
	myVmid = strconv.Itoa(int(vmid))
}

func GetSubscriptTags() ([]data.Uptag, error) {
	url := "https://api.bilibili.com/x/relation/tags?vmid=" + myVmid
	var jsonData data.TagRespBody
	err := data.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func GetSubscriptUpByTag(tagid int) ([]data.Up, error) {
	url := "https://api.bilibili.com/x/relation/tag?tagid=" + strconv.Itoa(tagid)
	var jsonData data.UpRespBody
	err := data.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func GetUp() ([]data.Up, error) {
	t, err := GetSubscriptTags()
	if err != nil {
		return nil, fmt.Errorf("error get tags: %s", err)
	}
	var up []data.Up
	for _, tag := range t {
		u, err := GetSubscriptUpByTag(tag.Id)
		if err != nil {
			return nil, fmt.Errorf("error get up by tag: %s", err)
		}
		up = append(up, u...)
	}
	return up, nil
}

func GetSubSeason() ([]data.Season, error) {
	url := "https://api.bilibili.com/x/v3/fav/folder/collected/list?pn=1&ps=20&up_mid=" + myVmid + "&platform=web"
	var jsonData data.SubSeasonRespBody
	err := data.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get subseason from %s: %s", url, err)
	}
	var season []data.Season
	for _, s := range jsonData.Data.List {
		season = append(season, data.Season{
			UpID:     s.UpID,
			SeasonID: s.SeasonID,
			Name:     s.Name,
			Label:    data.NoLab,
		})
	}
	return season, nil
}

func AllVg() (vgFilter []data.Vgroup, err error) { // get all vgroup from mysql
	// get up
	var ups []data.Up
	var vg []data.Vgroup
	err = dbSql.Db.Find(&ups).Error
	if err != nil {
		return nil, fmt.Errorf("get up fail: %v", err)
	}
	// log.Printf("get up success, %#v", ups)
	for i := range ups {
		vg = append(vg, &ups[i])
	}

	// get season
	var seasons []data.Season
	err = dbSql.Db.Find(&seasons).Error
	if err != nil {
		return nil, fmt.Errorf("get season fail: %v", err)
	}
	// log.Printf("get season success, %#v", seasons)
	for i := range seasons {
		vg = append(vg, &seasons[i])
	}

	// filter vg for label >= 3
	for _, v := range vg {
		if v.GetLabel() >= 3 {
			vgFilter = append(vgFilter, v)
		}
	}

	return
}
