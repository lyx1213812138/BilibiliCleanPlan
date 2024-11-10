package recommend

import (
	"log"
	"sort"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	dbsql "github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/spf13/viper"
)

type VideoSlice []data.Video

func (vs VideoSlice) Len() int {
	return len(vs)
}

func (vs VideoSlice) Less(i, j int) bool {
	// 优先级：标签 > 播放量
	v1, v2 := vs[i], vs[j]
	if v1.Label != v2.Label {
		return v1.Label > v2.Label
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
	// TODO: 查找已看过视频

	// fliter vg
	var allVs []data.Video
	for _, v := range vg {
		var vs []data.Video
		err := dbsql.Db.Where("up_id = ?", v.GetUpID()).Limit(LimitOfEachVg).Find(&vs).Error
		if err != nil {
			log.Printf("error get video by up: %s\n", err)
		}
		allVs = append(allVs, vs...)
	}

	// fliter seen
	res := VideoSlice{}
	for _, v := range allVs {
		if !dbsql.IfSeen(v.Bvid) {
			res = append(res, v)
		}
	}
	sort.Sort(res)
	// fmt.Println("allres: ", res)
	return res[:min(numRecommond, len(res))], nil
}

// TODO: 同一个up的视频，不要出现太多次，用参数而不是判断
