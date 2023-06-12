package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/cache"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/model"
	"strings"
)

func Recovery() gin.HandlerFunc {
	// 捕获panic
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		// 继续执行
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	// 获取请求头中的key
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
		// 判断key是否在缓存中
		if v, ok := cache.Cache.Get(cache.Key(key)); !ok || v == nil {
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

func Ctx() gin.HandlerFunc {
	// 为每个请求创建一个上下文
	return func(c *gin.Context) {
		// 创建一个上下文
		ctx := context.Background()
		ctx = context.WithValue(ctx, model.HeaderKey, c.GetHeader(model.HeaderKey))
		// 将上下文放入gin上下文中
		c.Set(model.GinCtxCtx, ctx)
		// 继续执行
		c.Next()
	}
}

func Error() gin.HandlerFunc {
	// 处理错误
	return func(c *gin.Context) {
		c.Next()
		// 如果有错误
		if len(c.Errors) > 0 {
			// 返回错误
			c.AbortWithStatusJSON(200, model.GeneralRes{Msg: c.Errors[0].Error()})
			return
		}
	}
}
