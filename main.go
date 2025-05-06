package main

import (
	"fmt"
	"net/http"
)

func main() {

	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Kiss me!")
	})

	//
	println(">> Listening to :4000...")
	_ = http.ListenAndServe(":4000", nil)
}
