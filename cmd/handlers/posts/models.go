package postsHandler

import (
	"github.com/Solar-2020/GoUtils/context"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	"github.com/valyala/fasthttp"
)

type postService interface {
	Create(ctx context.Context, request models.InputPost) (response models.Post, err error)
	GetList(ctx context.Context, request models.GetPostListRequest) (response []models.PostResult, err error)
}

type postTransport interface {
	CreateDecode(ctx *fasthttp.RequestCtx) (request models.InputPost, err error)
	CreateEncode(response models.Post, ctx *fasthttp.RequestCtx) (err error)

	GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error)
	GetListEncode(response []models.PostResult, ctx *fasthttp.RequestCtx) (err error)
}

type errorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) (err error)
	ServeFatalError(ctx *fasthttp.RequestCtx)
}
