//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package model

import (
	"errors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
)

const HeaderKey = "key"
const GinCtxCtx = "ctx"

type GeneralRes struct {
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

type VerifyInterface interface {
	Verify() error
}

func Verify[T any](req *T) error {
	if v, ok := any(req).(VerifyInterface); ok {
		return v.Verify()
	}
	return nil
}

type SdpPostReq struct {
	KeyTTL int    `json:"key_ttl"`                  // key的过期时间，单位秒，最大值为30
	Data   string `json:"data" validate:"required"` // sdp序列化为json字符串
}

func (s *SdpPostReq) Verify() error {
	if s.Data == "" {
		return errors.New(custerrors.DataIsEmpty)
	}
	return nil
}
