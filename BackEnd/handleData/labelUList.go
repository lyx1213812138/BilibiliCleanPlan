package handleData

import (
	"context"
	"fmt"

	"github.com/lyx1213812138/BilibiliCleanPlan/myMongodb"
	"github.com/lyx1213812138/BilibiliCleanPlan/myType"
)

func Label() {
	// aaa := myType.Up{Name: "aaa", UpID: 1, Label: myType.Ignore}
	// fmt.Println(aaa)
	// return
	var data []myType.DetailEachTag
	datap := &data
	err := myMongodb.Find("UpList", nil, datap)
	if err != nil {
		panic(fmt.Errorf("error find data from mongodb: %s", err))
	}
	// fmt.Println(data)
LabelLoop:
	for tagid, v := range data {
		for upid, _ := range v.UpList {
			upPtr := &data[tagid].UpList[upid]
			if (*upPtr).Label == myType.NoLab {
				fmt.Printf("UpID: %d, Name: %s, Label: %d\n", (*upPtr).UpID, (*upPtr).Name, (*upPtr).Label)
				fmt.Print("Input label (Break for -1): ")
				var label int
				fmt.Scan(&label)
				if label < 1 || label > 4 {
					fmt.Println("break")
					break LabelLoop
				}
				(*upPtr).Label = myType.Label(label)
			}
		}
	}
	// fmt.Println(data)
	fmt.Print("If save data to mongodb? (y/n): ")
	var save string
	fmt.Scan(&save)
	if save != "y" {
		return
	}
	myMongodb.Database.Drop(context.Background())
	err = myMongodb.Insert("UpList", data)
	if err != nil {
		panic(fmt.Errorf("error insert data to mongodb: %s", err))
	} else {
		fmt.Println("insert successfully")
	}
}
