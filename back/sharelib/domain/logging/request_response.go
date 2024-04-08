package logging

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

// bufferedResponseWriter レスポンスをバッファするResponseWriter実装
// レスポンスをログ出力するため、バッファーに一時的に蓄積する。
type bufferedResponseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func (b *bufferedResponseWriter) Write(data []byte) (int, error) {
	b.buffer.Write(data)
	return b.ResponseWriter.Write(data)
}

func (b *bufferedResponseWriter) String() string {
	return b.buffer.String()
}

// LogRequestResponseMiddleware リクエスト・レスポンスをログ出力するミドルウェア
func LogRequestResponseMiddleware(c *gin.Context) {
	// リクエスト内容をログ出力する
	body, _ := io.ReadAll(c.Request.Body)
	Debug("request",
		Var("method", c.Request.Method),
		Var("uri", c.Request.RequestURI),
		Var("remoteAddr", c.Request.RemoteAddr),
		Var("headers", c.Request.Header),
		Var("body", string(body)),
	)

	// HTTPリクエストのボディはストリームであり、一度読み取ると再度読み取ることはできない。
	// 再度利用する可能性があるので、リクエストボディを再設定する。
	if len(body) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	// レスポンス内容をログ出力するため、ResponseWriterをwrapする
	buffered := bufferedResponseWriter{
		ResponseWriter: c.Writer,
		buffer:         new(bytes.Buffer),
	}
	c.Writer = &buffered

	// 次の処理
	c.Next()

	// レスポンス内容をログ出力する
	Debug("response",
		Var("statusCode", buffered.Status()),
		Var("body", buffered.String()),
	)
}
