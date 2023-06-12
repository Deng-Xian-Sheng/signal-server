package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/middleware"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/service"
)

func Router(router *gin.Engine) {
	// 使用 gin-swagger 中间件 /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root := router.Group("/")
	{
		root.Use(middleware.Recovery(), middleware.Auth(), middleware.Ctx(), middleware.Error())
		offer := root.Group("offer/")
		{
			sdp := offer.Group("sdp/")
			{
				// GET /offer/sdp
				sdp.GET("", service.Offer.SdpGet)
				// POST /offer/sdp
				sdp.POST("", service.Offer.SdpPost)
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
