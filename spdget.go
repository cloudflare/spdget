package main

import (
    "fmt"
    "github.com/jmckaskill/gospdy"
    "net/http"
    "log"
    "flag"
    "io/ioutil"
)

func main() {
    follow := flag.Bool("L", false, "follow redirects")
    headers := flag.Bool("i", false, "show headers")
    flag.Parse()
    url := flag.Arg(0)
    full_url := "https://" + url
    tr := &spdy.Transport{}

    var err error
    var r *http.Response
    if *follow {
        client := &http.Client{Transport: tr}
        r, err = client.Get(full_url)
    } else {
        /* https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/PmYm3h2J3FE */
        var req *http.Request
        req, err = http.NewRequest("GET", full_url,nil)
        r, err = tr.RoundTrip(req)
    }
    if err != nil {
        log.Fatal(err)
        return
    }

    if *headers {
        for k := range r.Header {
            fmt.Println(k, r.Header[k])
        }
        fmt.Println()
    }

    body, err := ioutil.ReadAll(r.Body)
    fmt.Printf(string(body))
    defer r.Body.Close()

}
