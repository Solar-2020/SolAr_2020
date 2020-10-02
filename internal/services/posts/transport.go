package posts

import (
	"encoding/json"
	"github.com/BarniBl/SolAr_2020/internal/models"
	"github.com/valyala/fasthttp"
	"time"
)

type Transport interface {
	CreateDecode(ctx *fasthttp.RequestCtx) (response models.InputPost, err error)
	CreateEncode(response models.Post, ctx *fasthttp.RequestCtx) (err error)

	GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error)
	GetListEncode(response []models.Post, ctx *fasthttp.RequestCtx) (err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) CreateDecode(ctx *fasthttp.RequestCtx) (request models.InputPost, err error) {
	//userID := ctx.Value("UserID").(int)
	userID := 1
	var inputPost models.InputPost
	err = json.Unmarshal(ctx.Request.Body(), &inputPost)
	if err != nil {
		return
	}
	inputPost.PublishDate = time.Now()
	inputPost.CreateBy = userID
	request = inputPost
	return
}

func (t transport) CreateEncode(response models.Post, ctx *fasthttp.RequestCtx) (err error) {
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}

func (t transport) GetListDecode(ctx *fasthttp.RequestCtx) (request models.GetPostListRequest, err error) {
	//userID := ctx.Value("UserID").(int)
	request.UserID = ctx.Value("UserID").(int)
	request.GroupID = ctx.QueryArgs().GetUintOrZero("groupID")
	request.Limit = ctx.QueryArgs().GetUintOrZero("limit")

	startFrom := string(ctx.QueryArgs().Peek("startFrom"))
	request.StartFrom, err = time.Parse("2006-01-02", startFrom)
	if err != nil {
		return
	}

	return
}

func (t transport) GetListEncode(response []models.Post, ctx *fasthttp.RequestCtx) (err error) {
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}
