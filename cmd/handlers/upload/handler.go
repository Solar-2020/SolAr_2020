package uploadHandler

import (
	"github.com/Solar-2020/GoUtils/context"
)

type Handler interface {
	File(ctx context.Context)
	Photo(ctx context.Context)
}

type handler struct {
	uploadService   uploadService
	uploadTransport uploadTransport
	errorWorker     errorWorker
}

func NewHandler(uploadService uploadService, uploadTransport uploadTransport, errorWorker errorWorker) Handler {
	return &handler{
		uploadService:   uploadService,
		uploadTransport: uploadTransport,
		errorWorker:     errorWorker,
	}
}

func (h *handler) File(ctx context.Context) {
	request, err := h.uploadTransport.FileDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	response, err := h.uploadService.File(request)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.uploadTransport.FileEncode(response, ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) Photo(ctx context.Context) {
	request, err := h.uploadTransport.PhotoDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	response, err := h.uploadService.Photo(request)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.uploadTransport.PhotoEncode(response, ctx.RequestCtx)
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
