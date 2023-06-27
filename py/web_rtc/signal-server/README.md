# signal-server 

HTTP-based high concurrency signaling server for signaling exchange of Web RTC.（github.com/pion/webrtc/examples/pion-to-pion is her main user）

### Swagger API Docs

Gen Docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

If you are running signal-server, you can view the interface documentation by accessing the http://localhost:8080/swagger/index.html.

### Usage

Here is a supplement to the swagger interface documentation, which is obviously difficult to get started with the swagger interface documentation alone.



[//]: # (TODO )

```go
package custerrors // import "make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"


CONSTANTS

const (
	KeyIsEmpty                       = "key为空"
	KeyIsTooLong                     = "key太长，长度不能超过32位"
	KeyNotFound                      = "key未找到"
	CtxIsEmpty                       = "ctx为空"
	CacheOverrunMaxCost              = "缓存超出MaxCost，服务已达到最大负载"
	CacheOnEvictNoGetKey             = "缓存OnEvict回调，无法获取原始key"
	QueueNewFailBecauseAlreadyExists = "队列创建失败，因为已经存在"
	QueueGetFailBecauseNotExists     = "队列获取失败，因为不存在"
	DataIsEmpty                      = "data为空"
	KeyIsExpired                     = "key已过期"

	// offer 或 answer 的 sdp
	SdpNoValues     = "sdp没有值" // 对方（offer或answer）未放入，或已经被取出
	SdpAlreadyExist = "sdp已存在" // 对方（offer或answer）已经放入过一次，且只能放入一次

	// offer 或 answer 的 candidate
	CandidateNoValues = "candidate没有值" // 对方（offer或answer）从未放入，或已全部取出

	BodyIsEmpty = "body为空"
)
```
