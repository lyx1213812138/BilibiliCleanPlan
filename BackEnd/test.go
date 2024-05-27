package main

import (
	"context"

	store "github.com/lyx1213812138/BilibiliCleanPlan/store"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
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
	}
}
