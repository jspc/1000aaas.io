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

func get_a(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, a.String())
}

func proxy (c web.C, w http.ResponseWriter, r *http.Request) {
   switch c.URLParams["proxy"] {
    default:
        url = "http://jelly.stupids.org:6565"
    }

    req, err := http.NewRequest("Get", url, nil)
    if err != nil {
        response = err.Error()
    } else {
        defer req.Body.Close()
        body, err := ioutil.ReadAll(req.Body)
        if err != nil {
            response = err.Error()
        } else {
            response = string(body[:])
        }
    }

    fmt.Fprint(w, response)
}

func healthcheck(c web.C, w http.ResponseWriter, r *http.Request) {
    name, _ := os.Hostname()
    fmt.Fprintf(w, "Healthy! host: %s", name)
}

func main() {
    for i := 0; i < 1000; i++ {
        a.WriteString("a")
    }

    goji.Get("/", healthcheck)
    goji.Get("/v1/a", get_a)
    goji.Get("/v1/proxy/:proxy", proxy)
    goji.Serve()
}
