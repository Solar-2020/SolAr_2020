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
}

func NewHandler(postService postService, postTransport postTransport) Handler {
	return &handler{
		postService:   postService,
		postTransport: postTransport,
	}
}

func (h *handler) Create(ctx *fasthttp.RequestCtx) {
	post, err := h.postTransport.CreateDecode(ctx)
	if err != nil {

	}

	postReturn, err := h.postService.Create(post)
	if err != nil {

	}

	err = h.postTransport.CreateEncode(postReturn, ctx)
	if err != nil {

	}
}

func (h *handler) GetList(ctx *fasthttp.RequestCtx) {
	getPostListRequest, err := h.postTransport.GetListDecode(ctx)
	if err != nil {

	}

	postReturn, err := h.postService.GetList(getPostListRequest)
	if err != nil {

	}

	err = h.postTransport.GetListEncode(postReturn, ctx)
	if err != nil {

	}
}
