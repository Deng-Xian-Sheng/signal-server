package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/model"
)

func GetKey(ctx context.Context) string {
	if v, ok := ctx.Value("key").(string); ok {
		return v
	}
	log.Panicln(custerrors.KeyIsEmpty)
	return ""
}

func GetCtx(c *gin.Context) context.Context {
	if v, ok := c.Get(model.GinCtxCtx); ok {
		if ctx, ok := v.(context.Context); ok {
			return ctx
		}
	}
	log.Panicln(custerrors.CtxIsEmpty)
	return nil
}
