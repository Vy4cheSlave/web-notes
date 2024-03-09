package main

import (
	"html/template"
	"net/http"
	"os"
)

//// (1)
// import (
// 	"strconv"
// )

//// (2)
// import (
// 	"encoding/json"
// 	"io"
// )

var templates = template.Must(template.ParseFiles("index.html"))

func generalHandle(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		rootPathHandle(w, r)
	case "POST":
		rootPathHandlePost(w, r)
	}
}

func rootPathHandle(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r)
	templates.Execute(w, nil)
}

type File struct {
	Tilte string `json:"title"`
	Body  string `json:"body"`
}

func rootPathHandlePost(w http.ResponseWriter, r *http.Request) {

	fileName := "aboba.txt"
	resultFile, err := os.Open(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode("Unable to open file ")
		return
	}
	defer resultFile.Close()

	// //// (1)
	// // force a download with the content- disposition field
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fileName))
	// w.Header().Set("Content-Type", "application/octet-stream")
	// w.WriteHeader(http.StatusOK)
	// // serve file out.
	// http.ServeFile(w, r, fileName)
	// ////

	// //// (2)
	// data, _ := io.ReadAll(resultFile)
	// jsonFile := File{
	// 	Tilte: fileName,
	// 	Body:  string(data),
	// }
	// w.Header().Set("Content-type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(jsonFile)
	// ////
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", generalHandle)

	http.ListenAndServe("127.0.0.1:3001", mux)
}
