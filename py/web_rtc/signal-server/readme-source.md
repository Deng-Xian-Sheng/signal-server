[//]: # (自述文件生成命令)
[//]: # (p="README.md" && cat readme-source.md > ${p} && echo '\n\n```go' >> ${p} && go doc -all -src custerrors/custerrors.go >> ${p} && echo '```' >> ${p} -->)
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

[//]: # (TODO)