[//]: # (自述文件生成命令)
[//]: # (p="README.md" && cat readme-source.md > ${p} && echo '\n\n```go' >> ${p} && go doc -all -src custerrors/custerrors.go >> ${p} && echo '```' >> ${p} -->)
# signal-server 

HTTP-based high concurrency signaling server for signaling exchange of Web RTC.（[github.com/pion/webrtc/examples/pion-to-pion](https://github.com/pion/webrtc/examples/pion-to-pion) is her main user）

### Swagger API Docs

Gen Docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

If you are running signal-server, you can view the interface documentation by accessing the [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

### Usage

Here is a supplement to the swagger interface documentation, which is obviously difficult to get started with the swagger interface documentation alone.

Firstly, let's take a look at the common return parameters.

It looks like this:

```json
{
    "msg":"",
    "data":null
}
```

Please note that when msg is "", the request is processed normally.
The "normal" here means that the request has produced a result that meets your expectations.

When the msg is not "", an exception occurred in the request, which may not be severe and cannot be called an error.

If an exception occurs, the requester needs to handle it, and there are some defined msgs at the end of the readme file.
Please note that undefined msg may also be generated.

Let's start using this signaling server, TODO

### FAQ

- Inexplicable logs

If you see an error like this in the log:

```log
2023/07/02 19:47:32 队列获取失败，因为不存在
2023/07/02 19:47:32 队列获取失败，因为不存在
2023/07/02 19:47:32 队列获取失败，因为不存在
2023/07/02 19:47:32 cache OnReject回调发生错误  队列获取失败，因为不存在

2023/07/02 19:47:32 cache OnReject回调发生错误  队列获取失败，因为不存在

2023/07/02 19:47:32 cache OnReject回调发生错误  队列获取失败，因为不存在

2023/07/02 19:47:32 队列获取失败，因为不存在
2023/07/02 19:47:32 cache OnReject回调发生错误  队列获取失败，因为不存在
```

Please don't panic! This is not a big problem, because the "process" is not completed in its entirety, the requester does not seem to be using the signaling service correctly, and everything is a problem of the requester.

This happens when the service recycles the cache and queue, and if you want to know the details, read the source code.

- CacheOverrunMaxCost

If you see an error like this in the log:

```log
2023/07/02 19:47:32 缓存超出MaxCost，服务已达到最大负载
```

This means that the cache of the service has reached its maximum load, which results in the current inability to receive new SDP creation requests, but existing SDPs are not affected.

To explain, creating an SDP is the beginning of a "process", and not being able to create an SDP means that a new "process" cannot be created, but an existing "process" is not affected.

The "process" is the process by which an answer and offer establish a WebRTC connection through a signaling service.

Only when the cache is reclaimed can it continue to receive new SDP creation requests.

If your hardware configuration is high enough, you can try increasing the value of Max Cost, it is not in the configuration file, you need to modify the code, it is in the init() function of the cache package.

Be sure to read the documentation for the ristretto package before making changes [github.com/dgraph-io/ristretto](https://github.com/dgraph-io/ristretto)

**It should be noted that the DefMaxCacheTTL constant does not strictly determine the cache time of the key, and if the requester continues to generate valid requests, the key will always exist.**

**Obviously, this is a security issue, I haven't figured out how to solve it, if you have a good idea, welcome to bring it up.**

```go
package custerrors // import "signal-server/custerrors"


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
