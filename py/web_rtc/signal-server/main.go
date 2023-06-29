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

	"make_data_set_so-vits-svc/py/web_rtc/signal-server/router"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	gin.DefaultWriter = os.Stderr
	r := gin.Default()
	r.Use(cors.Default())
	router.Router(r)
	if err := r.Run(ipPort); err != nil {
		log.Panicln(err)
	}
}
