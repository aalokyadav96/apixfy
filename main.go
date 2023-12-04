package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",Index)
	http.HandleFunc("/v1",API)
	http.HandleFunc("/v1/gfycats/fetch/status/",Status)
	http.HandleFunc("/v1test",Test)
	
	//~ fs := http.FileServer(http.Dir("./uploads"))
	//~ http.Handle("/pix.gif", http.StripPrefix("/", fs))

	http.ListenAndServe(":8365",nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w,"V1")
}

func API(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w,"V1")
}

func Test(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w,"V1")
}


func Status(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println(r.URL.Path)
	FileName := r.URL.Path[len("/v1/gfycats/fetch/status/"):]
	var jsonStr = "{\"task\":\"complete\",\"gfyname\":\"" + FileName + "\"}"
	fmt.Println(jsonStr)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, jsonStr)
}

func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "https://xfy.onrender.com")
}
