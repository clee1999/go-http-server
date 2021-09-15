package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ROUTE : GET /
func hourHandler(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	h := currentTime.Hour()
	min := currentTime.Minute()
	fmt.Fprintf(w, "%v%v%v", h, "h", min)

}

// Save data to file
func saveData(entry, author string) {

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(author + ":" + entry + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}

}

// ROUTE : POST / ADD
func addHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	author := req.Form.Get("author")
	entry := req.Form.Get("entry")
	fmt.Fprintf(w, "%v:%v", author, entry)
	saveData(entry, author)

}

// ROUTE : GET / ENTRIES
func entriesHandler(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Panicf("impossible read file: %s", err)
	}
	fmt.Printf("%s", data)
	fmt.Fprintf(w, "%s", data)
}

func main() {
	http.HandleFunc("/", hourHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
