package handleData

import (
	"fmt"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	dbsql "github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
)

func Label() {
	ups := []data.Up{}
	err := dbsql.Db.Find(&ups).Error
	if err != nil {
		panic(err)
	}
	for upid := range ups {
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
	err = dbsql.Db.Save(&ups).Error
	if err != nil {
		panic(fmt.Errorf("error insert data to mongodb: %s", err))
	} else {
		fmt.Println("insert successfully")
	}
}
