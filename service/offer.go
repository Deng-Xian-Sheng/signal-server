//Copyright (c) [2023] [JinCanQi]
//[signal-server] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package service

import (
	"io"
	"log"
	"signal-server/cache"
	"signal-server/common"
	"signal-server/custerrors"
	"signal-server/model"
	"signal-server/queue"
	"strconv"
	"strings"
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
	ctx := common.Ctx.GetCtx(c)

	ttl, ok := cache.Cache.Get(common.Ctx.GetKey(ctx))
	if !ok || ttl == nil {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsExpired})
		return
	}
	ttlInt64, err := strconv.Atoi(ttl.(string))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	cache.Cache.SetWithTTL(common.Ctx.GetKey(ctx), "", 1, time.Duration(ttlInt64))

	answerQueue := queue.GetAnswerSdpQueue(common.Ctx.GetKey(ctx))

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
	ctx := common.Ctx.GetCtx(c)

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

	cache.Cache.SetWithTTL(common.Ctx.GetKey(ctx), "", 1, time.Second*time.Duration(req.KeyTTL))

	if queue.HasOfferSdpQueue(common.Ctx.GetKey(ctx)) {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.SdpAlreadyExist})
		return
	}
	queue.NewOfferSdpQueue(common.Ctx.GetKey(ctx)).Push(req.Data)

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
	ctx := common.Ctx.GetCtx(c)

	ttl, ok := cache.Cache.Get(common.Ctx.GetKey(ctx))
	if !ok || ttl == nil {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsExpired})
		return
	}
	ttlInt64, err := strconv.Atoi(ttl.(string))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	cache.Cache.SetWithTTL(common.Ctx.GetKey(ctx), "", 1, time.Duration(ttlInt64))

	candidateQueue := queue.GetAnswerCandidateQueue(common.Ctx.GetKey(ctx))

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
// @Security ApiKeyAuth
// @Router /offer/candidate [post]
func (o *offer) CandidatePost(c *gin.Context) {
	ctx := common.Ctx.GetCtx(c)

	ttl, ok := cache.Cache.Get(common.Ctx.GetKey(ctx))
	if !ok || ttl == nil {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsExpired})
		return
	}
	ttlInt64, err := strconv.Atoi(ttl.(string))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}

	cache.Cache.SetWithTTL(common.Ctx.GetKey(ctx), "", 1, time.Duration(ttlInt64))

	var req string
	reqByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: err.Error()})
		return
	}
	req = string(reqByte)

	if strings.TrimSpace(req) == "" {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.BodyIsEmpty})
		return
	}

	if queue.HasOfferCandidateQueue(common.Ctx.GetKey(ctx)) {
		queue.GetOfferCandidateQueue(common.Ctx.GetKey(ctx)).Push(req)
	} else {
		queue.NewOfferCandidateQueue(common.Ctx.GetKey(ctx)).Push(req)
	}

	c.JSON(200, model.GeneralRes{})
}
