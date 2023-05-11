package loghelper

import (
	"fmt"
	"net"
	"strings"
)

//获取MAC
func GetIP() string {

	//var ips []string
	//
	//ips = getIPs()

	//return strings.Join(ips, "|")
	return "本机"
}

func getIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)

		if ipNet, isValidIpNet = address.(*net.IPNet); isValidIpNet && !ipNet.IP.IsLoopback() {

			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.String()
				if strings.Index(ip, "169.") == -1 {
					ips = append(ips, ip)
				}
			}
		}
	}
	return ips
}
