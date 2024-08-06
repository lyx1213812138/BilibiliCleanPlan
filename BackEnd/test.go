package main

import (
	"fmt"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	_ "github.com/lyx1213812138/BilibiliCleanPlan/myMongodb"
	_ "github.com/lyx1213812138/BilibiliCleanPlan/recommend"
)

func main() {
	// vg, err := myMongodb.GetAllVgroup()
	// if err != nil {
	// 	panic(err)
	// }
	// vs, err := recommend.RecommondList(vg)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(vs)
	s, err := data.GetSubSeason()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
