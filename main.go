package main

import (
    "os"
    "bytes"
    "fmt"
    "net/http"
    "strconv"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var a bytes.Buffer
var response string
var url string
var count_int int

func get_root(c web.C, w http.ResponseWriter, r *http.Request) {
    name, _ := os.Hostname()
    fmt.Fprintf(w, "Healthy! host: %s", name)
}

func get_a(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, a.String())
}

func get_cock(c web.C, w http.ResponseWriter, r *http.Request) {
    cock := cock(20)
    fmt.Fprintf(w, cock)
}

func get_cock_with_count(c web.C, w http.ResponseWriter, r *http.Request) {
    count_int, err := strconv.Atoi(c.URLParams["count"])

    if err != nil {
        http.Error(w, "Pass me a number, dick", 400)
    }

    cock := cock(count_int)
    fmt.Fprintf(w, cock)
}

func cock(count int) string{
    var cock bytes.Buffer
    cock.WriteString("8")
    for i := 0; i < count; i++ {
        cock.WriteString("=")
    }
    cock.WriteString("D")
    return cock.String()
}

func main() {
    for i := 0; i < 1000; i++ {
        a.WriteString("a")
    }

    goji.Get("/", get_root)
    goji.Get("/v1/a", get_a)
    goji.Get("/v1/cock", get_cock)
    goji.Get("/v1/cock/:count", get_cock_with_count)

    /* Financial Times specific API test magic */
    goji.Get("/1000aaas/v1/a", get_a)
    goji.Serve()
}
