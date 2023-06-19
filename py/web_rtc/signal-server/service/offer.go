//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package service

import (
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/cache"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/model"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/queue"
	"time"

	"github.com/gin-gonic/gin"
)

type offer struct{}

var Offer = &offer{}

// SdpGet
// @Summary 获取 answer 发送给 offer 的 sdp
// @Description 获取 answer 发送给 offer 的 sdp
// @Tags offer
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /offer/sdp [get]
func (o *offer) SdpGet(c *gin.Context) {
	ctx := GetCtx(c)

	answerQueue := queue.GetAnswerSdpQueue(GetKey(ctx))
	
	if answerQueue.Len() == 0 {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.SdpNoValues})
		return
	}

	c.JSON(200, model.GeneralRes{Data: answerQueue.Pop().(string)})
}

// SdpPost
// @Summary offer 发送 sdp 给 answer
// @Description offer 发送 sdp 给 answer
// @Tags offer
// @Accept json
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Param body body model.SdpPostReq true "请求参数，具体见结构体"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /offer/sdp [post]
func (o *offer) SdpPost(c *gin.Context) {
	ctx := GetCtx(c)

	var req model.SdpPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	if err := model.Verify(&req); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	cache.Cache.SetWithTTL(GetKey(ctx), "", 1, time.Second*time.Duration(req.KeyTTL))

	queue.NewOfferSdpQueue(GetKey(ctx)).Push(req.Data)

	c.JSON(200, model.GeneralRes{})
}

// CandidateGet
// @Summary 获取 answer 发送给 offer 的 candidate
// @Description 获取 answer 发送给 offer 的 candidate
// @Tags offer
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /offer/candidate [get]
func (o *offer) CandidateGet(c *gin.Context) {
	ctx := GetCtx(c)

	candidateQueue := queue.GetAnswerCandidateQueue(GetKey(ctx))

	if candidateQueue.Len() == 0 {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.CandidateNoValues})
		return
	}

	c.JSON(200, model.GeneralRes{Data: candidateQueue.Pop().(string)})
}

// CandidatePost
// @Summary offer 发送 candidate 给 answer
// @Description offer 发送 candidate 给 answer
// @Tags offer
// @Accept text/plain
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Param body body string true "请求参数，candidate字符串"
// @Success 200 {object} model.GeneralRes
func (o *offer) CandidatePost(c *gin.Context) {
	ctx := GetCtx(c)

	var req string
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	ttl, ok := cache.Cache.GetTTL(GetKey(ctx))
	if !ok {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsExpired})
		return
	}

	cache.Cache.SetWithTTL(GetKey(ctx), "", 1, ttl)

	queue.NewOfferCandidateQueue(GetKey(ctx)).Push(req)

	c.JSON(200, model.GeneralRes{})
}
