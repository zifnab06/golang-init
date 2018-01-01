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

func setupInterfaces(config Config) {
    log.SetOutput(os.Stderr)
    printInterfaces()
    for _, iface := range config.Network {
        log.Println("Initializing ", iface.Name, "...")
        if link, err := netlink.LinkByName(iface.Name); err == nil {
            if addr, err := netlink.ParseAddr(iface.Address); err == nil {
                if err := netlink.AddrAdd(link, addr); err != nil {
                    log.Println("Error: unable to assign address to eth0", err)
                }
                if err := netlink.LinkSetUp(link); err != nil {
                    log.Println("Error: unable to bring ", iface.Name, " up", err)
                }
                printInterfaces()
                if len(iface.Gateway) > 0 {
                    log.Println("Adding default route to ", iface.Name, "...")
                    gw := net.ParseIP(iface.Gateway)
                    route := netlink.Route{LinkIndex: link.Attrs().Index, Gw: gw}
                    if err:= netlink.RouteAdd(&route); err != nil {
                        log.Println("Error: unable to add route", err)
                    }
                }
            } else {
                log.Fatal("Error: unable to parse address")
            }
        } else {
            log.Fatal("Error: unable to get eth0")
        }
    }
}
