package recommend

import (
	"log"
	"sort"
	"time"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	dbsql "github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/lyx1213812138/BilibiliCleanPlan/get"
	"github.com/spf13/viper"
)

type VideoSlice []data.Video

func filtSeenVideo(allVs []data.Video) ([]data.Video, error) {
	log.Print("after time", time.Unix((time.Now().Unix()-viper.GetInt64("min_create_time_before_now")), 0))
	res := VideoSlice{}
	for _, v := range allVs {
		if !get.IfSeen(v.Bvid) && v.CreateTime > int(time.Now().Unix())-viper.GetInt("min_create_time_before_now") {
			res = append(res, v)
		}
	}
	return res, nil
}

func (vs VideoSlice) Len() int {
	return len(vs)
}

func (vs VideoSlice) Less(i, j int) bool {
	// 优先级：标签 > 播放量
	v1, v2 := vs[i], vs[j]
	now := int(time.Now().Unix())
	t1, t2 := now-v1.CreateTime, now-v2.CreateTime
	tt1, tt2 := t1/3e6, t2/3e6 // 大概一个月
	// TODO: 优化推荐
	if v1.Label != v2.Label {
		return v1.Label > v2.Label
	} else if tt1 != tt2 {
		return tt1 < tt2
	} else {
		return v1.View > v2.View
	}
}

func (vs VideoSlice) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func RecommondList(vg []data.Vgroup) ([]data.Video, error) {
	numRecommond := viper.GetInt("recommend.num")
	LimitOfEachVg := max(3, numRecommond/len(vg)*2)

	// fliter vg (limit)
	var allVs []data.Video
	for _, v := range vg {
		var vs []data.Video
		err := dbsql.Db.Where("up_id = ?", v.GetUpID()).Limit(LimitOfEachVg).Find(&vs).Error
		if err != nil {
			log.Printf("error get video by up: %s\n", err)
		}
		allVs = append(allVs, vs...)
	}
	vsUnseen, err := filtSeenVideo(allVs)
	if err != nil {
		return nil, err
	}
	res := VideoSlice(vsUnseen)
	sort.Sort(res)
	return res[:min(numRecommond, len(res))], nil
}

func SearchList(vg []data.Vgroup, search string) ([]data.Video, error) {
	log.Printf("search: %s\n", search)
	// fliter vg (unlimit)
	var allVs []data.Video
	for _, v := range vg {
		var vs []data.Video
		err := dbsql.Db.Where("up_id = ?", v.GetUpID()).Find(&vs).Error
		if err != nil {
			log.Printf("error get video by up: %s\n", err)
		}
		allVs = append(allVs, vs...)
	}

	// fliter search
	res := VideoSlice{}
	for _, v := range allVs {
		if fuzzy.Match(search, v.Title) || fuzzy.Match(search, v.UpName) {
			res = append(res, v)
		}
	}

	sort.Sort(res)
	return res, nil
}
