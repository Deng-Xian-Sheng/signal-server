//Copyright (c) [2023] [JinCanQi]
//[signal-server] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package tool

import (
	"net"
)

// GetNetWorkIp 获取当前设备上所有的 IPv4 地址。
// 它通过调用 net.Interfaces() 函数获取了当前设备上所有的网络接口，然后遍历每个网络接口，获取该接口上的所有 IP 地址。
// 最后，它将所有非回环地址的 IPv4 地址返回。
// 我在本机测试，获得了
// 169.254.78.149
// 192.168.1.1
// 我感觉很玄学，如果你只需要一个ip地址，推荐选最后一个。（169.254.78.149是操作系统赋予的无法从外部访问的地址，大多用于操作系统内部）
func GetNetWorkIp() (ips []string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, i := range ifaces {
		addrs, err2 := i.Addrs()
		err = err2
		if err != nil {
			return
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
				ips = append(ips, ip.String())
			}
		}
	}
	return
}
