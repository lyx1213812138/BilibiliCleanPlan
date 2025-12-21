package get

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
)

func GetAndStoreHistory() error {
	var cursor data.Hcursor
	var historys []data.History
	cursor.Ps = 20

	// get last view at for limiting get span
	var lastViewAtBody []data.LastViewAt
	var lastViewAt int
	err := dbSql.Db.Find(&lastViewAtBody).Error
	if err != nil || len(lastViewAtBody) == 0 {
		log.Printf("error get last view at: %s\n", err)
		lastViewAt = int(time.Now().AddDate(0, 0, -7).Unix())
	} else {
		lastViewAt = lastViewAtBody[0].Value
	}
	log.Print("get history after last view at: ", time.Unix(int64(lastViewAt), 0))

	// get history after last view at
	end := false
	for !end {
		urlQuery := url.Values{}
		urlQuery.Add("max", strconv.Itoa(cursor.Max))
		urlQuery.Add("view_at", strconv.Itoa(cursor.View_at))
		urlQuery.Add("business", cursor.Business)
		urlQuery.Add("ps", strconv.Itoa(cursor.Ps))
		url := "https://api.bilibili.com/x/web-interface/history/cursor?" + urlQuery.Encode()
		var jsonData data.HistoryRespBody
		err := data.Get(url, &jsonData)
		if err != nil {
			return fmt.Errorf("error get history: %s", err)
		}
		cursor = jsonData.Data.Cursor
		for _, h := range jsonData.Data.List {
			historys = append(historys, h.History)
			if h.View_at < lastViewAt {
				end = true
				break
			}
		}
		if cursor.Max == 0 {
			end = true
		}
	}

	// store history
	err = dbSql.Db.Save(&historys).Error
	if err != nil {
		return fmt.Errorf("store history fail: %v", err)
	}

	// store last view at
	err = dbSql.Db.Save(&data.LastViewAt{Id: 1, Value: int(time.Now().Unix())}).Error
	if err != nil {
		return fmt.Errorf("store last view at fail: %v", err)
	}
	return nil
}
