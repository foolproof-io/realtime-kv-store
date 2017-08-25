package main

import (
        "bytes"
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
)

var kvs map[string]string

func init() {
    kvs = map[string]string{}
    host, _ := os.Hostname()
    kvs["host"] = host
}

func main() {
        server := http.NewServeMux()
        server.HandleFunc("/", handleKeyValue)

        log.Println("Server listening on port 80")
        log.Fatal(http.ListenAndServe(":80", server))
}

func handleKeyValue(w http.ResponseWriter, r *http.Request) {
    log.Printf("Serving request: %v", r)
    switch r.Method {
    case http.MethodGet:
        key := r.URL.Path[1:]
        value, ok := kvs[key]
        if !ok {
            http.Error(w, "no such key", http.StatusNotFound)
            return
        }
        fmt.Fprintln(w, value)

    case http.MethodPut:
        key := r.URL.Path[1:]
        buf := new(bytes.Buffer)
        _, err := io.Copy(buf, r.Body)
        if err != nil {
            log.Fatalf("error while reading request body: %s", err)
        }
        kvs[key] = buf.String()

    default:
        log.Printf("no idea what to do with %v\n", r)
        http.Error(w, "unsupported operation", http.StatusMethodNotAllowed)
    }
}
