package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/dbSql"
	"github.com/lyx1213812138/BilibiliCleanPlan/recommend"
)

var port string = "12121"

func server() {
	http.HandleFunc("/getvideo", handlerVideo)
	http.HandleFunc("/allvgroup", handlerVg)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Printf("server start at localhost:%s\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

// /getvideo
func handlerVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("request for url: ", r.URL.Path)
	// get request
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// make vgroup
	var req req
	var vg []data.Vgroup
	err = json.Unmarshal(b, &req)
	if err != nil {
		log.Printf("error unmarshal body: %s\nfind all\n", err)
		vg, err = dbsql.AllVg()
	} else if len(req.List) == 0 {
		log.Println("no data\nfind all")
		vg, err = dbsql.AllVg()
	} else {
		// log.Printf("request data: %#v\n", req)
		vg = reqToVg(req)
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("error get all vgroup from mongodb: ", err)
		return
	}

	vs, err := recommend.RecommondList(vg)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	b, err = json.Marshal(vs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}

func reqToVg(req req) (vg []data.Vgroup) {
	// fmt.Printf("req: %#v\n", req)
	for _, ele := range req.List {
		switch ele.Type {
		case data.IsSeason:
			vg = append(vg, &data.Season{
				UpID:     ele.UpId,
				SeasonID: ele.SeasonId,
				Label:    ele.Label,
			})
		case data.IsUp:
			vg = append(vg, &data.Up{
				UpID:  ele.UpId,
				Label: ele.Label,
			})
		}
	}
	return
}

type req struct {
	List []reqEle `json:"list"`
}

type reqEle struct {
	Type     data.VgType `json:"type"`
	UpId     int         `json:"mid"`
	SeasonId int         `json:"sid"`
	Label    data.Label  `json:"label"`
}

// /allvgroup
func handlerVg(w http.ResponseWriter, r *http.Request) {
	log.Println("request for url: ", r.URL.Path)
	// get request
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	vg, err := dbsql.AllVg()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("error get all vgroup from mongodb: ", err)
		return
	}
	var res string = "["
	// JSON解析器在解析时会忽略空白字符，包括空格、制表符、换行符等。
	for _, v := range vg {
		s, err := v.GetStr()
		if err != nil {
			fmt.Println("error get string from one vgroup: ", err)
			continue
		}
		res += s + ","
	}
	// remove last ',' to ']'
	if len(res) > 1 {
		res = res[:len(res)-1]
		res += "]"
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(res))
}
