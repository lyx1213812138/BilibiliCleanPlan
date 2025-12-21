package get

import (
	"fmt"
	"log"
	"time"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/spf13/viper"
)

func StoreAllVideo() error { // store all video to mysql
	vg, err := AllVg()
	if err != nil {
		return fmt.Errorf("get all vgroup fail: %v", err)
	}

	video, err := GetVideoByVg(vg)
	// var seasons []data.Season
	// err := myMongodb.Find("Seasons", nil, &seasons)
	if err != nil {
		return fmt.Errorf("get video fail: %v", err)
	}
	if video == nil {
		log.Printf("get video warning: video is nil")
		return nil
	}
	// log.Printf("get video success, %#v", video)
	res := dbSql.Db.Save(&video)
	// primary key may duplicate, so not use Create
	if res.Error != nil {
		return fmt.Errorf("store video fail: %v", res.Error)
	}
	log.Printf("mysql store video success, rows affected: %d", res.RowsAffected)
	return nil
}

func GetVideoByVg(vg []data.Vgroup) ([]data.Video, error) {
	// var wg sync.WaitGroup
	var videos []data.Video
	// var mu sync.Mutex
	for _, v := range vg {
		if v.GetLabel() <= 3 { // ignore or nolab
			continue
		}
		// wg.Add(1)
		// go func(v data.Vgroup) {
		// 	defer wg.Done()
		// 	video, err := v.GetVideo()
		// 	if err != nil {
		// 		fmt.Printf("error get video: %s\n", err)
		// 		return
		// 	}
		// 	mu.Lock()
		// 	videos = append(videos, video...)
		// 	mu.Unlock()
		// 	// fmt.Println("len video(utilFromDb) ", len(video))
		// }(v)
		time.Sleep(time.Duration(viper.GetInt("wait_time_ms")) * time.Millisecond) // 为了防止请求过快，导致被封ip
		video, err := v.GetVideo()
		if err != nil {
			fmt.Printf("error get video: %s\n", err)
		}
		videos = append(videos, video...)
	}
	// wg.Wait()
	fmt.Println(len(videos))
	return videos, nil
}

func IfSeen(bvid string) bool {
	var h data.History
	err := dbSql.Db.Take(&h, "bvid = ?", bvid).Error
	if err != nil {
		return false
	}
	return true
}
