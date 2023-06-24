//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package main

import (
	"flag"
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/docs"

	"make_data_set_so-vits-svc/py/web_rtc/signal-server/router"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/tool"
	"net"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 大致想法
// 写一个信令服务器
// 它不使用数据库也不使用redis 它支持并发 它通过内存存储数据（考虑使用库） 并且每条数据都有一个过期时间 过期后自动删除 过期时间最大30秒 可由请求方指定
// 它使用gin启动一个http服务，此服务只有一个接口，请求方式为post，

// @title Signal Server API
// @version 1.0
// @description 基于HTTP的高并发信令服务器，用于WebRTC的信令交换。（github.com/pion/webrtc/examples/pion-to-pion是她的主要用户）
// @host localhost:8080
// @BasePath /
func main() {
	log.SetOutput(os.Stderr)

	var ipPort string
	flag.StringVar(&ipPort, "it", ":8080", "ip:port")
	flag.Parse()

	if ipPort == "" {
		log.Panicln("请指定ip和port。")
	}

	if false {
		if host, port, err := net.SplitHostPort(ipPort); err != nil {
			log.Panicln(err)
		} else {
			if port == "" {
				log.Panicln("请指定port。")
			}
			if host == "" {
				ips, err := tool.GetNetWorkIp()
				if err != nil {
					log.Panicln(err)
				}
				if len(ips) > 0 {
					docs.SwaggerInfo.Host = net.JoinHostPort(ips[len(ips)-1], port)
				}
			}
		}
	}

	gin.DefaultWriter = os.Stderr
	r := gin.Default()
	r.Use(cors.Default())
	router.Router(r)
	if err := r.Run(ipPort); err != nil {
		log.Panicln(err)
	}
}
