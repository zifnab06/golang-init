package main

import (
        "log"
        "io/ioutil"
        "gopkg.in/yaml.v2"
)

type iface struct {
    Name string
    Address string
    Gateway string
}
type Config struct {
    Http struct {
        Port int
    }
    Ssh struct {
        Port int
        Keys map[string][]string
    }
    Network []iface
}

func LoadConfig() Config {
    config := Config{}
    yamlFile, err := ioutil.ReadFile("/etc/init.yml")
    if err != nil {
        log.Fatal(err)
    }
    if err := yaml.Unmarshal(yamlFile, &config); err != nil {
        log.Fatal(err)
    }
    log.Printf("%+v", config)
    return config
}
