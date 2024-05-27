package getData

import (
	"fmt"
	"strconv"

	"github.com/lyx1213812138/BilibiliCleanPlan/get"
)

func getSubscriptTags(vmid string) ([]uptag, error) { // vmid: 434100110
	url := "https://api.bilibili.com/x/relation/tags?vmid=" + vmid
	var jsonData tagRespBody
	err := get.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func getSubscriptUpByTag(tagid int) ([]up, error) {
	url := "https://api.bilibili.com/x/relation/tag?tagid=" + strconv.Itoa(tagid)
	var jsonData upRespBody
	err := get.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}
