package recommend

import (
	"fmt"
	"sort"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/myMongodb"
)

const (
	numRecommond = 20
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
	// TODO: 查找已看过视频
	allVs, err := myMongodb.GetVideoByVg(vg)
	if err != nil {
		return nil, fmt.Errorf("error get video by vgroup: %s", err)
	}
	res := VideoSlice{}
	for _, v := range allVs {
		if !v.Seen {
			res = append(res, v)
		}
	}
	sort.Sort(res)
	// fmt.Println("allres: ", res)
	return res[:min(numRecommond, len(res))], nil
}

// TODO: 同一个up的视频，不要出现太多次，用参数而不是判断
