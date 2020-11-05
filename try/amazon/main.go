package main

import (
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"

	_ "net/http/pprof"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var mux = httptreemux.NewContextMux()
	api := mux.NewGroup("/api")

	// sku-price 同步 统计
	api.GET("/amazon/account/credit", getCredits)

	if err := http.ListenAndServe(":10304", mux); err != nil {
		log.Fatal(err)
	}
}
