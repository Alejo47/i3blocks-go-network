package main

import (
	"flag"
	"fmt"
	"net"
)

var ifname = flag.String("if", "", "")
var up_icon = flag.String("up-icon", "+", "")
var down_icon = flag.String("down-icon", "-", "")
var up_color = flag.String("up-color", "#00FF00", "")
var down_color = flag.String("down-color", "#FF0000", "")

func init() {
	flag.Parse()
}

func main() {
	ifaces, _ := net.Interfaces()
	if len(*ifname) == 0 {
		
		return
	}
	for _, iface := range ifaces {
		if iface.Name == *ifname {
			if addrs, _ := iface.Addrs(); len(addrs) > 0 {
				fmt.Println(*up_icon)
				fmt.Println(*up_icon)
				fmt.Println(*up_color)
			} else {
				fmt.Println(*down_icon)
				fmt.Println(*down_icon)
				fmt.Println(*down_color)
			}
			return
		}
	}
	fmt.Println(*down_icon)
	fmt.Println(*down_icon)
	fmt.Println(*down_color)
}
