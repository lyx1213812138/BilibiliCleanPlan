// get data from bilibili api using get.go
package data

import (
	"fmt"
	"strconv"
	"sync"
)

var myVmid string

// LEARN : init函数在main函数执行之前，自动被调用
func SetMyVmid(vmid float64) {
	myVmid = strconv.Itoa(int(vmid))
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

func GetUp() ([]Up, error) {
	t, err := GetSubscriptTags()
	if err != nil {
		return nil, fmt.Errorf("error get tags: %s", err)
	}
	var up []Up
	for _, tag := range t {
		u, err := GetSubscriptUpByTag(tag.Id)
		if err != nil {
			return nil, fmt.Errorf("error get up by tag: %s", err)
		}
		up = append(up, u...)
	}
	return up, nil
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

func GetVideoByVg(vg []Vgroup) ([]Video, error) {
	var wg sync.WaitGroup
	var videos []Video
	var mu sync.Mutex
	for _, v := range vg {
		if v.GetLabel() <= 1 { // ignore or nolab
			continue
		}
		wg.Add(1)
		go func(v Vgroup) {
			defer wg.Done()
			video, err := v.GetVideo()
			if err != nil {
				fmt.Printf("error get video: %s\n", err)
				return
			}
			mu.Lock()
			videos = append(videos, video...)
			mu.Unlock()
			// fmt.Println("len video(utilFromDb) ", len(video))
		}(v)
	}
	wg.Wait()
	fmt.Println(len(videos))
	return videos, nil
}
