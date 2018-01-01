package main

import (
        "log"
        "os"
        "time"
)

func main() {
    log.SetOutput(os.Stderr)
    log.Println("Starting golang-init...")
    config := LoadConfig()
    setupInterfaces(config)
    go http_server(config)
    go ssh_server(config)
    select {}
}
