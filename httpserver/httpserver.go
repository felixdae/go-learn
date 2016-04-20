package main

import "net/http"
import "io"
import "log"

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main(){
    http.HandleFunc("/hello", HelloServer)
    log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
