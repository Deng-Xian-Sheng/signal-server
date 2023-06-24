package common

import (
	"context"
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/model"

	"github.com/gin-gonic/gin"
)

type ctx2 struct{

}

var Ctx = &ctx2{}

func (c2 *ctx2) GetCtx(c *gin.Context) context.Context {
	if v, ok := c.Get(model.GinCtxCtx); ok {
		if ctx, ok := v.(context.Context); ok {
			return ctx
		}
	}
	log.Panicln(custerrors.CtxIsEmpty)
	return nil
}

func (c2 *ctx2) GetKey(ctx context.Context) string {
	if v, ok := ctx.Value(model.HeaderKey).(string); ok {
		return v
	}
	log.Panicln(custerrors.KeyIsEmpty)
	return ""
}