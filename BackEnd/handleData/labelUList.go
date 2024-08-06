package handleData

import (
	"context"
	"fmt"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/myMongodb"
)

func Label() {
	ups := []data.Up{}
	err := myMongodb.Find("Ups", nil, &ups)
	if err != nil {
		panic(err)
	}
	for upid, _ := range ups {
		upPtr := &ups[upid]
		if (*upPtr).Label == data.NoLab {
			fmt.Printf("UpID: %d, Name: %s, Label: %d\n", (*upPtr).UpID, (*upPtr).Name, (*upPtr).Label)
			fmt.Print("Input label (Break for -1): ")
			var label int
			fmt.Scan(&label)
			if label < 1 || label > 4 {
				fmt.Println("break")
				break
			}
			(*upPtr).Label = data.Label(label)
		}
	}
	// fmt.Println(data)
	fmt.Print("If save data to mongodb? (y/n): ")
	var save string
	fmt.Scan(&save)
	if save != "y" {
		return
	}
	myMongodb.Database.Collection("Ups").Drop(context.TODO())
	err = myMongodb.Insert("Ups", ups)
	if err != nil {
		panic(fmt.Errorf("error insert data to mongodb: %s", err))
	} else {
		fmt.Println("insert successfully")
	}
}
