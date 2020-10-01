package posts

import (
	"encoding/json"
	"github.com/BarniBl/SolAr_2020/internal/models"
	"github.com/valyala/fasthttp"
)

type Transport interface {
	CreateDecode(ctx *fasthttp.RequestCtx) (response models.InputPost, err error)
	CreateEncode(response models.InputPost, ctx *fasthttp.RequestCtx) (err error)

	GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error)
	GetListEncode(response []models.InputPost, ctx *fasthttp.RequestCtx) (err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) CreateDecode(ctx *fasthttp.RequestCtx) (response models.InputPost, err error) {
	userID := ctx.Value("UserID").(int)

	if err = json.Unmarshal(ctx.Request.Body(), &response); err != nil {
		return
	}

	response.CreateBy = userID
	return
}

func (t transport) CreateEncode(response models.InputPost, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}

func (t transport) GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error) {
	panic("implement me")
}

func (t transport) GetListEncode(response []models.InputPost, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}
