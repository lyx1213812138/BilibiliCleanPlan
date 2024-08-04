package getData

import (
	"fmt"
	"strconv"

	"github.com/lyx1213812138/BilibiliCleanPlan/get"

	"github.com/lyx1213812138/BilibiliCleanPlan/myType"
)

func getSubscriptTags(vmid string) ([]myType.Uptag, error) { // vmid: 434100110
	url := "https://api.bilibili.com/x/relation/tags?vmid=" + vmid
	var jsonData myType.TagRespBody
	err := get.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

func getSubscriptUpByTag(tagid int) ([]myType.Up, error) {
	url := "https://api.bilibili.com/x/relation/tag?tagid=" + strconv.Itoa(tagid)
	var jsonData myType.UpRespBody
	err := get.Get(url, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error get from %s: %s", url, err)
	}
	return jsonData.Data, nil
}

/*func storeUpList() {
	// tags, _ := getSubscriptTags("434100110")
	// store.Insert("upsTag", tags)
	cur, err := store.Database.Collection("upsTag").Find(context.TODO(), options.Find())
	if err != nil {
		panic(err)
	}
	var upstag []uptag
	err = cur.All(context.TODO(), &upstag)
	if err != nil {
		panic(err)
	}
	store.Database.Collection("UpList").Drop(context.TODO())
	// fmt.Printf("%#v", upstag)
	for _, v := range upstag {
		var dataEachTag detailEachTag
		dataEachTag.TagInfo = v
		dataEachTag.UpList, err = getSubscriptUpByTag(v.Tagid)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%#v\n\n\n\n\n", dataEachTag)
		err = store.Insert("UpList", []any{dataEachTag})
		if err != nil {
			panic(err)
		}

}*/
