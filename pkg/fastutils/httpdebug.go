package fastutils

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/tsingson/fastws-example/pkg/vtils"
)

// RequestCtxDebug log all requestCtx field
func RequestCtxDebug(ctx *fasthttp.RequestCtx, log *zap.Logger, payloadDetail bool) {
	log.Info(" RequestCtxDebug ==================================================================================")

	uri := ctx.RequestURI()

	log.Info("RequestCtxDebug",
		zap.String("uri", vtils.B2S(uri)))

	args := ctx.QueryArgs()

	if args.Len() > 0 {
		args.VisitAll(
			func(key, value []byte) {
				log.Info("args",
					zap.String("key", vtils.B2S(key)),
					zap.String("value", vtils.B2S(value)))
			})
	}

	ctx.Request.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		log.Info("header",
			zap.String("key", vtils.B2S(key)),
			zap.String("value", vtils.B2S(value)))
	})

	log.Info("payload size",
		zap.Int("payload", len(ctx.Request.Body())))
	if payloadDetail {
		log.Info("payload",
			zap.String("payload", vtils.B2S(ctx.Request.Body())))
	}
}

// RequestDebug  log request all field
func RequestDebug(req *fasthttp.Request, log *zap.Logger, payloadDetail bool) {
	log.Info(" RequestDebug ==================================================================================")

	req.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		log.Info("header",
			zap.String("key", vtils.B2S(key)),
			zap.String("value", vtils.B2S(value)))
	})

	log.Warn("payload size",
		zap.Int("payload", len(req.Body())))
	if payloadDetail {
		log.Info("payload",
			zap.String("payload", vtils.B2S(req.Body())))
	}
}

// ResponseDebug  log response all field
func ResponseDebug(resp *fasthttp.Response, log *zap.Logger, payloadDetail bool) {
	log.Info(" ResponseDebug ==================================================================================")

	log.Info("http status code",
		zap.Int("code", resp.StatusCode()))

	resp.Header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		log.Info("header",
			zap.String("key", vtils.B2S(key)),
			zap.String("value", vtils.B2S(value)))
	})

	log.Info("payload size",
		zap.Int("payload", len(resp.Body())))
	if payloadDetail {
		log.Info("payload",
			zap.String("payload", vtils.B2S(resp.Body())))
	}
}
