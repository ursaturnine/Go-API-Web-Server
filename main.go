package main

import (
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

func main() {
	// main route
	http.HandleFunc("/", rootHandler)
	// start a server
	err := http.ListenAndServe("localhost:8080", nil)
	// end program if error
	if err != nil {
		panic(err)
		os.Exit(1)
	}

}
