package myMongodb

import (
	"context"
	"fmt"
	"sync"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
)

/* 弃用 func AllUp() ([]myType.Up, error) {
	var data []myType.DetailEachTag
	err := Find("UpList", nil, &data)
	if err != nil {
		return nil, fmt.Errorf("error find data from mongodb: %s", err)
	}
	var up []myType.Up
	for _, v := range data {
		up = append(up, v.UpList...)
	}
	return up, nil
} */

func RmSameUps() {
	var updata, res []data.Up
	err := Find("Ups", nil, &updata)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(updata))
	var mapp = make(map[int]struct {
		l  data.Label
		id int
	})
	for _, v := range updata {
		if _, ok := mapp[v.UpID]; !ok {
			mapp[v.UpID] = struct {
				l  data.Label
				id int
			}{l: v.Label, id: len(res)}
			res = append(res, v)
		} else if m := mapp[v.UpID]; m.l < v.Label {
			res[m.id].Label = v.Label
			m.l = v.Label
			mapp[v.UpID] = m
		}
	}
	fmt.Println(len(res))
	Database.Collection("Ups").Drop(context.TODO())
	Insert("Ups", res)
}

/* 添加标签 */
func AddLabel() {
	err := Insert("Ups", []data.Up{{405054588, "TED精选演讲", 4}})
	if err != nil {
		panic(err)
	}
	RmSameUps()
}

func GetAllVgroup() ([]data.Vgroup, error) {
	var ups []data.Up
	var seasons []data.Season
	var vgroups []data.Vgroup
	err := Find("Ups", nil, &ups)
	if err != nil {
		return nil, fmt.Errorf("error find ups from mongodb: %s", err)
	}
	err = Find("Seasons", nil, &seasons)
	if err != nil {
		return nil, fmt.Errorf("error find seasons from mongodb: %s", err)
	}
	// merge ups and seasons into vgroups
	for _, v := range ups {
		vgroups = append(vgroups, data.Vgroup(&v))
	}
	for _, v := range seasons {
		vgroups = append(vgroups, data.Vgroup(&v))
	}
	return vgroups, nil
}

func GetVideoByVg(vg []data.Vgroup) ([]data.Video, error) {
	var wg sync.WaitGroup
	var videos []data.Video
	var mu sync.Mutex
	for _, v := range vg {
		if v.GetLabel() <= 1 { // ignore or nolab
			continue
		}
		wg.Add(1)
		go func(v data.Vgroup) {
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
