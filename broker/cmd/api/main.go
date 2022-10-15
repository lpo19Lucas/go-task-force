package main

import (
    "fmt"
    "log"
    "net/http"
)

const PORT = "80"

type Config struct {

}   

func main() {
    app := &Config{}

    log.Println("Starting broker service on port ...", PORT)

    s := &http.Server{
        Addr:           fmt.Sprintf(":%s", PORT),
        Handler:        app.routes(),
        // ReadTimeout:    10 * time.Second,
        // WriteTimeout:   10 * time.Second,
        // MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
    // err != nil {
    //     log.Panic(err),
    // }
}