package upload

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
	"github.com/valyala/fasthttp"
)

type Transport interface {
	FileDecode(ctx *fasthttp.RequestCtx) (request models.WriteFile, err error)
	FileEncode(response models.File, ctx *fasthttp.RequestCtx) (err error)

	PhotoDecode(ctx *fasthttp.RequestCtx) (request models.WritePhoto, err error)
	PhotoEncode(response models.Photo, ctx *fasthttp.RequestCtx) (err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) FileDecode(ctx *fasthttp.RequestCtx) (request models.WriteFile, err error) {
	panic("implement me")
}

func (t transport) FileEncode(response models.File, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}

func (t transport) PhotoDecode(ctx *fasthttp.RequestCtx) (request models.WritePhoto, err error) {
	panic("implement me")
}

func (t transport) PhotoEncode(response models.Photo, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}




