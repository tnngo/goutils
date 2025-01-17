package goutils

import (
	"net"
)

type LocalNetworkInfo struct {
	IP4        string
	IP6        string
	MacAddress string
	IfaceName  string
}

func GetLocalNetworkInfo() ([]*LocalNetworkInfo, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	localNetworkInfos := make([]*LocalNetworkInfo, 0)
	for _, iface := range interfaces {
		localNetworkInfo := &LocalNetworkInfo{
			IfaceName:  iface.Name,
			MacAddress: iface.HardwareAddr.String(),
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		var isLoopback bool
		for _, addr := range addrs {
			// Filter out loopback addresses
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					localNetworkInfo.IP4 = ipnet.IP.String()
				}
				if ipnet.IP.To16() != nil && ipnet.IP.To4() == nil {
					localNetworkInfo.IP6 = ipnet.IP.String()
				}
			} else {
				isLoopback = true
			}
		}

		if isLoopback {
			continue
		}

		localNetworkInfos = append(localNetworkInfos, localNetworkInfo)
	}

	return localNetworkInfos, nil
}
