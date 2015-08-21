package main

import (
    "os"
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var a bytes.Buffer
var response string
var url string

func get_root(c web.C, w http.ResponseWriter, r *http.Request) {
    name, _ := os.Hostname()
    fmt.Fprintf(w, "Healthy! host: %s", name)
}

func get_a(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, a.String())
}

func main() {
    for i := 0; i < 1000; i++ {
        a.WriteString("a")
    }

    goji.Get("/", get_root)
    goji.Get("/v1/a", get_a)
    goji.Serve()
}
