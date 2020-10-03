package postsHandler

import (
	"github.com/valyala/fasthttp"
)

type Handler interface {
	Create(ctx *fasthttp.RequestCtx)
	GetList(ctx *fasthttp.RequestCtx)
}

type handler struct {
	postService   postService
	postTransport postTransport
	errorWorker   errorWorker
}

func NewHandler(postService postService, postTransport postTransport, errorWorker errorWorker) Handler {
	return &handler{
		postService:   postService,
		postTransport: postTransport,
		errorWorker:   errorWorker,
	}
}

func (h *handler) Create(ctx *fasthttp.RequestCtx) {
	post, err := h.postTransport.CreateDecode(ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	postReturn, err := h.postService.Create(post)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.postTransport.CreateEncode(postReturn, ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}

func (h *handler) GetList(ctx *fasthttp.RequestCtx) {
	getPostListRequest, err := h.postTransport.GetListDecode(ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	postReturn, err := h.postService.GetList(getPostListRequest)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.postTransport.GetListEncode(postReturn, ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}
