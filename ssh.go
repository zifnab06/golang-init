package main

import (
    "fmt"
    "io"
    "log"
    "strings"
    "github.com/gliderlabs/ssh"
    gossh "golang.org/x/crypto/ssh"
)

func get_key(key string) string {
    return strings.Split(key, " ")[1]
}

func check_key(keys []string, key string) bool {
    for _, k := range keys {
        k := get_key(k)
        if k == key {
            return true
        }
    }
    return false
}

func ssh_server(config Config) {
    ssh.Handle(func(s ssh.Session) {
        io.WriteString(s, fmt.Sprintf("Welcome %s\n", s.User()))
    })

    publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
        for _, k := range config.Ssh.Keys[ctx.User()] {
            pk, _, _, _, err := gossh.ParseAuthorizedKey([]byte(k))
            if err != nil {
                return false;
            }
            if ssh.KeysEqual(pk, key) {
                return true;
            }
        }
        return false
    })

    log.Println(fmt.Sprintf("Starting ssh server on port %v", config.Ssh.Port))
    log.Fatal(ssh.ListenAndServe(fmt.Sprintf(":%v", config.Ssh.Port), nil, publicKeyOption))
}
