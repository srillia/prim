/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 17:27
 */

package helper

import (
	"net"
	"os"
)

// 获取服务器Ip
func GetServerIp() string {
	var (
		ip       string
		hostname string
	)
	hostname = os.Getenv("PRIM_HOSTNAME")
	if hostname != "" {
		return hostname
	}
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}
	return ip
}
