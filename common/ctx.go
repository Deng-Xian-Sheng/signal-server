//Copyright (c) [2023] [JinCanQi]
//[signal-server] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package common

import (
	"context"
	"log"
	"signal-server/custerrors"
	"signal-server/model"

	"github.com/gin-gonic/gin"
)

type ctx2 struct {
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
