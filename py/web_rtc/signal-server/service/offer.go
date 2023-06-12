package service

import (
	"github.com/gin-gonic/gin"
)

type offer struct{}

var Offer = &offer{}

func (o *offer) SdpGet(c *gin.Context) {

}

// SdpPost
// @Summary offer 发送 sdp 给 answer
// @Description offer 发送 sdp 给 answer
// @Tags offer
// @Accept octet-stream
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Param body body string true "sdp"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /offer/sdp [post]
func (o *offer) SdpPost(c *gin.Context) {
	ctx := GetCtx(c)

}

func (o *offer) CandidateGet(c *gin.Context) {

}

func (o *offer) CandidatePost(c *gin.Context) {

}
