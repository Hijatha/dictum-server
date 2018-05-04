package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./" + r.URL.Path)
	fmt.Printf("%s\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8090", nil);err != nil {
        log.Fatal("ListenAndServe: ", err)
}
}
