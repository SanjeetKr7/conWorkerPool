package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	r.HandleFunc("/allComic", AllComic).Methods("GET")
}

func AllComic(w http.ResponseWriter, r *http.Request) {
	// func AllComic() {
	// https://xkcd.com/info.0.json
	noOfJobs, err := fetchTotalComic()
	if err != nil {
		log.Fatal(err)
	}
	go AllocateJobs(noOfJobs)

	// get results
	done := make(chan bool)
	go GetResults(done)

	// create worker pool
	noOfWorkers := 100
	CreateWorkerPool(noOfWorkers)

	// wait for all results to be collected
	<-done

	// convert result collection to JSON
	data, err := json.MarshalIndent(ResultCollection, "", "    ")
	if err != nil {
		log.Fatal("json err: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	// write json data to file
	err = writeToFile(data)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToFile(data []byte) error {
	f, err := os.Create("xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
