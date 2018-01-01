package main

import (
        "log"
        "net"
        "os"
        "github.com/vishvananda/netlink"
)

func printInterfaces() {
    if interfaces, err := net.Interfaces(); err == nil {
        for _, inter := range interfaces {
            if addrs, err := inter.Addrs(); err == nil {
                for _, addr := range addrs {
                    log.Println(inter.Name, "->", addr)
                }
            } else {
                log.Println(inter.Name, "->", err)
            }
        }
    } else {
        log.Println(err)
    }
}

func printRoutes() {
    if interfaces, err := net.Interfaces(); err == nil {
        for _, inter := range interfaces {
            if link, err := netlink.LinkByName(inter.Name); err == nil {
                if routes, err := netlink.RouteList(link, 4); err == nil {
                    for _, route := range routes {
                        log.Println("src:", route.Src, " dst:", route.Dst, " gw:", route.Gw)
                    }
                }
            }
        }
    }
}

func setupInterfaces() {
    log.SetOutput(os.Stderr)
    log.Println("Initializing eth0...")
    printInterfaces()
    if eth0, err := netlink.LinkByName("eth0"); err == nil {
        if addr, err := netlink.ParseAddr("173.255.208.61/24"); err == nil {
            if err := netlink.AddrAdd(eth0, addr); err != nil {
                log.Println("Error: unable to assign address to eth0", err)
            }
            if err := netlink.LinkSetUp(eth0); err != nil {
                log.Println("Error: unable to bring eth0 up", err)
            }
            printInterfaces()
            log.Println("Adding default route to eth0...")
            printRoutes()
            gw := net.IPv4(173, 255, 208, 1)
            route := netlink.Route{LinkIndex: eth0.Attrs().Index, Gw: gw}
            if err:= netlink.RouteAdd(&route); err != nil {
                log.Println("Error: unable to add route", err)
            }
            printRoutes()
        } else {
            log.Println("Error: unable to parse address")
        }
    } else {
        log.Println("Error: unable to get eth0")
    }
}

func main() {
    log.SetOutput(os.Stderr)
    log.Println("Starting golang-init...")
    config := LoadConfig()
    //setupInterfaces()
    log.Printf("Starting http echo server on port %v...", config.Http.Port)
    go http_server(config)
    go ssh_server(config)
    for {}
}
