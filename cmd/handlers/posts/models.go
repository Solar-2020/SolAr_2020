package postsHandler

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
	"github.com/valyala/fasthttp"
)

type postService interface {
	Create(request models.InputPost) (response models.InsertPost, err error)
	GetList(request models.GetPostListRequest) (response []models.InsertPost, err error)
}

type postTransport interface {
	CreateDecode(ctx *fasthttp.RequestCtx) (request models.InputPost, err error)
	CreateEncode(response models.InsertPost, ctx *fasthttp.RequestCtx) (err error)

	GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error)
	GetListEncode(response []models.InsertPost, ctx *fasthttp.RequestCtx) (err error)
}
