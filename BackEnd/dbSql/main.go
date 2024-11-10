package dbsql

import (
	"fmt"
	"log"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	// "github.com/lyx1213812138/BilibiliCleanPlan/myMongodb"
)

func Store() error {
	vg, err := AllVg()
	if err != nil {
		return fmt.Errorf("get all vgroup fail: %v", err)
	}

	video, err := data.GetVideoByVg(vg)
	// var seasons []data.Season
	// err := myMongodb.Find("Seasons", nil, &seasons)
	if err != nil {
		return fmt.Errorf("get video fail: %v", err)
	}
	// log.Printf("get video success, %#v", video)
	res := Db.Save(&video)
	// primary key may duplicate, so not use Create
	if res.Error != nil {
		return fmt.Errorf("store video fail: %v", res.Error)
	}
	log.Printf("mysql store video success, rows affected: %d", res.RowsAffected)
	return nil
}

func AllVg() (vg []data.Vgroup, err error) {
	// get up
	var ups []data.Up
	err = Db.Find(&ups).Error
	if err != nil {
		return nil, fmt.Errorf("get up fail: %v", err)
	}
	// log.Printf("get up success, %#v", ups)
	for i := range ups {
		vg = append(vg, &ups[i])
	}

	// get season
	var seasons []data.Season
	err = Db.Find(&seasons).Error
	if err != nil {
		return nil, fmt.Errorf("get season fail: %v", err)
	}
	// log.Printf("get season success, %#v", seasons)
	for i := range seasons {
		vg = append(vg, &seasons[i])
	}
	return
}

func IfSeen(bvid string) bool {
	err := Db.First(data.SeenVideo{Bvid: bvid}).Error
	if err != nil {
		return false
	}
	return true
}
