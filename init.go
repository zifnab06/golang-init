package main

import (
        "log"
        "os"
)

func main() {
    log.SetOutput(os.Stderr)
    log.Println("Starting golang-init...")
    config := LoadConfig()
    // if length of network hosts is more than 0
    setupInterfaces(config)
    // endif
    log.Printf("Starting http echo server on port %v...", config.Http.Port)
    go http_server(config)
    go ssh_server(config)
    for {}
}
