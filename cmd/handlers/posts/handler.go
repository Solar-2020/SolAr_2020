package postsHandler

import (
	"github.com/Solar-2020/GoUtils/context"
)

type Handler interface {
	Create(ctx context.Context)
	GetList(ctx context.Context)
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

func (h *handler) Create(ctx context.Context) {
	post, err := h.postTransport.CreateDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	postReturn, err := h.postService.Create(ctx, post)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.postTransport.CreateEncode(postReturn, ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) GetList(ctx context.Context) {
	getPostListRequest, err := h.postTransport.GetListDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	postReturn, err := h.postService.GetList(ctx, getPostListRequest)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.postTransport.GetListEncode(postReturn, ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) handleError(err error, ctx context.Context) {
	err = h.errorWorker.ServeJSONError(ctx.RequestCtx, err)
	if err != nil {
		h.errorWorker.ServeFatalError(ctx.RequestCtx)
	}
	return
}
