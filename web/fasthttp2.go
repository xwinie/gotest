package main

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func httpHandle(ctx *fasthttp.RequestCtx) {
	b := fasthttp.AcquireByteBuffer()
	b.B = append(b.B, "Hello "...)
	// 这里是编码过的 HTML 文本了，>strong 等
	b.B = fasthttp.AppendHTMLEscape(b.B, "<strong>World</strong>")
	defer fasthttp.ReleaseByteBuffer(b) // 记得释放

	ctx.Write(b.B)
}

func main() {
	// 使用 fasthttprouter 创建路由
	router := fasthttprouter.New()
	router.GET("/", httpHandle)
	if err := fasthttp.ListenAndServe("0.0.0.0:12345", router.Handler); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}