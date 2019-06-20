package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const port = "8081"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", handle)
	addr := "0.0.0.0:" + port
	log.Printf("listening to %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println("error:", err)
	}
}

var prPool = sync.Pool{
	New: func() interface{} {
		return new(pullRequest)
	},
}

type pullRequest struct {
	PullRequest struct{ ID int } `json:"pull_request"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	var data = prPool.New().(*pullRequest)
	defer prPool.Put(data)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println("err", err)
		http.Error(w, "internal server error ise8988", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("pull request id: %d", data.PullRequest.ID)))
}
