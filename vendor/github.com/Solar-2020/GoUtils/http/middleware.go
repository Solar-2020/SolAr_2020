package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Middleware interface {
	CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler
	Log(next fasthttp.RequestHandler) fasthttp.RequestHandler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return middleware{}
}

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func (m middleware) CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func (m middleware) Log(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Start new request: ", ctx.Request.URI())
		fmt.Println(ctx.Request.String())
		next(ctx)
	}
}

func NewLogCorsChain(middleware Middleware) func(func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
	return func(target func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
		return middleware.Log(middleware.CORS(target))
	}
}