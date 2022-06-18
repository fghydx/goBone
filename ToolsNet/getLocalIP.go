package ToolsNet

import "net"

func GetLocalIP() []string {
	var res []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = append(res, ipnet.IP.String())
			}
		}
	}
	return res
}
