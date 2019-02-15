package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port int

func main() {
	flag.IntVar(&port,"port", 8000, "port to run web server on")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"status": "up"}`))
		if err != nil {
			http.Error(w,"unable to write response", http.StatusInternalServerError)
			return
		}
		//w.WriteHeader(http.StatusOK)
		return
	})
	fmt.Printf("Starting on port: %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d",port),nil); err != nil {
		panic(err)
	}
}
