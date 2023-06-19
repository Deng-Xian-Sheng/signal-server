//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "make_data_set_so-vits-svc/py/web_rtc/signal-server/docs"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/middleware"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/service"
)

func Router(router *gin.Engine) {
	// 使用 gin-swagger 中间件 /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root := router.Group("/")
	root.Use(middleware.Ctx(), middleware.Error())
	{
		offer := root.Group("offer/")
		{
			sdp := offer.Group("sdp/")
			{
				// POST /offer/sdp
				sdp.POST("", service.Offer.SdpPost)
				root.Use(middleware.Auth())
				// GET /offer/sdp
				sdp.GET("", service.Offer.SdpGet)
			}
			candidate := offer.Group("candidate/")
			{
				// GET /offer/candidate
				candidate.GET("", service.Offer.CandidateGet)
				// POST /offer/candidate
				candidate.POST("", service.Offer.CandidatePost)
			}
		}
		answer := root.Group("answer/")
		{
			sdp := answer.Group("sdp/")
			{
				// GET /answer/sdp
				sdp.GET("", service.Answer.SdpGet)
				// POST /answer/sdp
				sdp.POST("", service.Answer.SdpPost)
			}
			candidate := answer.Group("candidate/")
			{
				// GET /answer/candidate
				candidate.GET("", service.Answer.CandidateGet)
				// POST /answer/candidate
				candidate.POST("", service.Answer.CandidatePost)
			}
		}
	}
}
