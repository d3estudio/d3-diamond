package main

import (
    "net/http"
    "log"
)

func main () {
    http.
    HandleFunc("/",
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Mundo!!"))
    })

    log.Fatal(http.ListenAndServe(":8000", nil))
}
