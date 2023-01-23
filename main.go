package main

import (
	"fmt"
	"net/http"
	"os"
)



func main(){
  http.HandleFunc("/",helloWorld)
  fmt.Println("Listening on port 8080")
  err := http.ListenAndServe(":8080",nil)
  fmt.Println(err)
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
  host, err := os.Hostname()
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Fprint(w,fmt.Sprintf("Hello k8s from %s\n",host))
}
