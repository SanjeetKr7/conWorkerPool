package main

import (
	"conWorkerPool/pkg"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// create a router from gorilla/mux
	r := mux.NewRouter()

	// register all the enpoints
	pkg.RegisterRouters(r)
	// pkg.AllComic()

	// start listing and serving http request
	http.ListenAndServe(":8081", r)
}
