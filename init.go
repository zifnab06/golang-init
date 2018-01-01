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
    log.Println("Sleep 5")
    time.Sleep(5)
    log.Println("End sleep")
    go http_server(config)
    go ssh_server(config)
    for {}
}
