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

type answer struct{}

var Answer = &answer{}

// SdpGet
// @Summary 获取 offer 发送给 answer 的 sdp
// @Description 获取 offer 发送给 answer 的 sdp
// @Tags answer
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /answer/sdp [get]
func (o *answer) SdpGet(c *gin.Context) {
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

	offerQueue := queue.GetOfferSdpQueue(common.Ctx.GetKey(ctx))

	if offerQueue.Len() == 0 {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.SdpNoValues})
		return
	}

	c.JSON(200, model.GeneralRes{Data: offerQueue.Pop().(string)})
}

// SdpPost
// @Summary answer 发送 sdp 给 offer
// @Description answer 发送 sdp 给 offer
// @Tags answer
// @Accept text/plain
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Param body body string true "请求参数，sdp序列化为json字符串"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /answer/sdp [post]
func (o *answer) SdpPost(c *gin.Context) {
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
	if queue.HasAnswerSdpQueue(common.Ctx.GetKey(ctx)) {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.SdpAlreadyExist})
		return
	}

	queue.NewAnswerSdpQueue(common.Ctx.GetKey(ctx)).Push(req)

	c.JSON(200, model.GeneralRes{})
}

// CandidateGet
// @Summary 获取 offer 发送给 answer 的 candidate
// @Description 获取 offer 发送给 answer 的 candidate
// @Tags answer
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /answer/candidate [get]
func (o *answer) CandidateGet(c *gin.Context) {
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

	offerQueue := queue.GetOfferCandidateQueue(common.Ctx.GetKey(ctx))

	if offerQueue.Len() == 0 {
		c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.CandidateNoValues})
		return
	}

	c.JSON(200, model.GeneralRes{Data: offerQueue.Pop().(string)})
}

// CandidatePost
// @Summary answer 发送 candidate 给 offer
// @Description answer 发送 candidate 给 offer
// @Tags answer
// @Accept text/plain
// @Produce json
// @Param key header string true "key是一个随机字符串由请求方生成，长度不能超过32位，用于鉴权"
// @Param body body string true "请求参数，candidate序列化为json字符串"
// @Success 200 {object} model.GeneralRes
// @Security ApiKeyAuth
// @Router /answer/candidate [post]
func (o *answer) CandidatePost(c *gin.Context) {
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

	if queue.HasAnswerCandidateQueue(common.Ctx.GetKey(ctx)) {
		queue.GetAnswerCandidateQueue(common.Ctx.GetKey(ctx)).Push(req)
	} else {
		queue.NewAnswerCandidateQueue(common.Ctx.GetKey(ctx)).Push(req)
	}

	c.JSON(200, model.GeneralRes{})
}
