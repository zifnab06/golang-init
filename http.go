package main

import (
        "log"
        "fmt"
        "net/http"
)

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
    log.Println(request.RemoteAddr + " requested " + request.URL.Path)
    request.Write(writer)
}

func http_server(config Config) {
    http.HandleFunc("/", EchoHandler)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Http.Port), nil))
}

