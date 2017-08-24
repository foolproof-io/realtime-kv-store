package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
        server := http.NewServeMux()
        server.HandleFunc("/", hello)
        log.Println("Server listening on port 80")
        log.Fatal(http.ListenAndServe(":80", server))
}

func hello(w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving request: %s", r.URL.Path)
        host, _ := os.Hostname()
        fmt.Fprintf(w, "Hello, world!\n")
        fmt.Fprintf(w, "Hostname: %s\n", host)
}
