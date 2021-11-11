package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/otiai10/gosseract"
)

func main(){

  fmt.Println(":8080 is ready_")

  r := mux.NewRouter()

  r.HandleFunc("/upload", Upload).Methods("POST")

  http.ListenAndServe(":8080", r)
}

func Upload(w http.ResponseWriter, r *http.Request){

  encoder := json.NewEncoder(w)

  file, handler, err := r.FormFile("myFile")

  if err != nil{

    fmt.Printf("%v with upload file", err)
    
  }

  defer file.Close()

  client := gosseract.NewClient()

  defer client.Close()

  client.SetLanguage("eng")

  client.SetImage(handler.Filename)

  text, _ := client.Text()

  fmt.Println(text)

  encoder.Encode(text)

}
