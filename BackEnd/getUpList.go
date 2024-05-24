package main

import (
	"fmt"
	"strconv"

	"github.com/lyx1213812138/BilibiliCleanPlan/get"
)

func getSubscriptTags(vmid string) ([]uptag, error) {
	url := "https://api.bilibili.com/x/relation/tags?vmid=" + vmid
	var jsonData tagRespBody
	err := get.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	if jsonData.Code != 0 {
		return nil, fmt.Errorf("error get from %s: %s", url, jsonData.Message)
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
	if jsonData.Code != 0 {
		return nil, fmt.Errorf("error get from %s: %s", url, jsonData.Message)
	}
	return jsonData.Data, nil
}

/*
特别关注：-10
短篇休闲：97606352
知识：97607696
长篇休闲：97616784
知识：97607696
*/
