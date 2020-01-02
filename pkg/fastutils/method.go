package fastutils

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// FastDebug fast http
type FastDebug struct {
	log           *zap.Logger
	payloadDetail bool
}

var defaultFastDebug = &FastDebug{
	payloadDetail: false,
}

// ZapLoggerOption options
type FastDebugOptions func(*FastDebug)

// WithDebug set debug
func WithDebug() FastDebugOptions {
	return func(o *FastDebug) {
		o.payloadDetail = true
	}
}

// WithDebug set debug
func WithLog(log *zap.Logger) FastDebugOptions {
	return func(o *FastDebug) {
		o.log = log
	}
}

// NewZapLogger  init a log
func New(opts ...FastDebugOptions) *FastDebug {
	s := defaultFastDebug
	for _, o := range opts {
		o(s)
	}
	return s
}

// Tid fast http connect id
func Tid(ctx *fasthttp.RequestCtx) string {
	tid := strconv.FormatInt(int64(ctx.ID()), 10)
	return tid
}

// RequestCtxDebug log all requestCtx field
func (s *FastDebug) RequestCtxDebug(ctx *fasthttp.RequestCtx) {
	fmt.Println(" RequestCtxDebug ==================================================================================")

	uri := ctx.RequestURI()

	s.log.Warn("RequestCtxDebug", zap.String("uri", vtils.B2S(uri)))

	args := ctx.QueryArgs()

	if args.Len() > 0 {
		args.VisitAll(
			func(key, value []byte) {
				s.log.Warn("args", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))
			})
	}

	ctx.Request.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		s.log.Warn("header", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))
	})

	s.log.Warn("payload size", zap.Int("payload", len(ctx.Request.Body())))
	if s.payloadDetail {
		s.log.Warn("payload", zap.String("payload", vtils.B2S(ctx.Request.Body())))
	}
}

// RequestDebug  log request all field
func (s *FastDebug) RequestDebug(req *fasthttp.Request) {
	fmt.Println(" RequestDebug ==================================================================================")
	s.log.Warn("RequestDebug")

	req.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		s.log.Warn("header", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))
	})

	s.log.Warn("payload size", zap.Int("payload", len(req.Body())))
	if s.payloadDetail {
		s.log.Warn("payload", zap.String("payload", vtils.B2S(req.Body())))
	}
}

// ResponseDebug  log response all field
func (s *FastDebug) ResponseDebug(resp *fasthttp.Response) {
	fmt.Println(" ResponseDebug ==================================================================================")
	s.log.Warn("ResponseDebug")

	resp.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		s.log.Warn("header", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))
	})

	s.log.Warn("payload size", zap.Int("payload", len(resp.Body())))
	if s.payloadDetail {
		s.log.Warn("payload", zap.String("payload", vtils.B2S(resp.Body())))
	}
}
