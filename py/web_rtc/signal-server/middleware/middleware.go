//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package middleware

import (
	"context"
	"errors"
	"fmt"

	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/cache"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/common"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// 中间件注册优先级 99
func Recovery() gin.HandlerFunc {
	// 捕获panic
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				c.Error(errors.New(fmt.Sprint(err)))
			}
		}()
		// 继续执行
		c.Next()
	}
}

// 中间件注册优先级 97
func Auth() gin.HandlerFunc {
	// 获取请求头中的key
	return func(c *gin.Context) {
		// 判断key是否在缓存中
		if v, ok := cache.Cache.Get(common.Ctx.GetKey(common.Ctx.GetCtx(c))); !ok || v == nil {
			// 如果key不在缓存中
			// 返回错误
			c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyNotFound})
			return
		}

		// 如果key正确
		// 继续执行
		c.Next()
	}
}

// 中间件注册优先级 98
func Ctx() gin.HandlerFunc {
	// 为每个请求创建一个上下文
	return func(c *gin.Context) {
		// 获取请求头中的key
		key := c.GetHeader(model.HeaderKey)
		key = strings.TrimSpace(key)
		// 如果key为空
		if key == "" {
			// 返回错误
			c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsEmpty})
			return
		}
		// 如果key不为空
		// 判断key的长度是否大于32位
		if len([]rune(key)) > 32 {
			// 返回错误
			c.AbortWithStatusJSON(200, model.GeneralRes{Msg: custerrors.KeyIsTooLong})
			return
		}
		// 创建一个上下文
		ctx := context.Background()
		ctx = context.WithValue(ctx, model.HeaderKey, key)
		// 将上下文放入gin上下文中
		c.Set(model.GinCtxCtx, ctx)
		// 继续执行
		c.Next()
	}
}

// 中间件注册优先级 100
func Error() gin.HandlerFunc {
	// 处理错误
	return func(c *gin.Context) {
		c.Next()
		// 如果有错误
		if len(c.Errors) > 0 {
			 c.Writer.Write([]byte{})
			// 返回错误
			l := make([]string, len(c.Errors))
			for i, v := range c.Errors {
				l[i] = v.Error()
			}
			c.AbortWithStatusJSON(200, model.GeneralRes{Msg: strings.Join(l, ",")})
			return
		}
	}
}
