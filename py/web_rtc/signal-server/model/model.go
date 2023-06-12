package model

const HeaderKey = "key"
const GinCtxCtx = "ctx"

type GeneralRes struct {
	Msg  string `json:"msg"`
	Data interface{}
}
