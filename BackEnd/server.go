package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lyx1213812138/BilibiliCleanPlan/data"
	"github.com/lyx1213812138/BilibiliCleanPlan/recommend"
)

var port string = "12121"

func Server() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// get request
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	req := new(req)
	err = json.Unmarshal(b, req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// make data
	var vg []data.Vgroup
	for _, ele := range req.List {
		switch ele.Type {
		case data.IsSeason:
			vg = append(vg, &data.Season{
				UpID:     ele.UpId,
				SeasonID: ele.SeasonId,
			})
		case data.IsUp:
			vg = append(vg, &data.Up{
				UpID: ele.UpId,
			})
		}
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
	w.Write(b)
}

type req struct {
	List []reqEle `json:"list"`
}

type reqEle struct {
	Type     data.VgType `json:"type"`
	UpId     int         `json:"mid"`
	SeasonId int         `json:"sid"`
}
