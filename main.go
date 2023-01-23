package main

import (
	"fmt"
	"net/http"
)



func main(){
  http.HandleFunc("/",helloWorld)
  fmt.Println("Listening on port 8080")
  err := http.ListenAndServe(":8080",nil)
  fmt.Println(err)
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w,"Hello k8s!\n")
}
