package upload

import (
	"encoding/json"
	"fmt"
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
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}
	fmt.Println(form)
	//json.Unmarshal([]byte(form.Value["body"]),&request)
	return
}

func (t transport) FileEncode(response models.File, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}

func (t transport) PhotoDecode(ctx *fasthttp.RequestCtx) (request models.WritePhoto, err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	var tempBody string
	body := form.Value["body"]
	for i, _ := range body {
		tempBody += body[i]
	}

	var tempFile string
	file := form.Value["file"]
	for i, _ := range file {
		tempFile += file[i]
	}

	err = json.Unmarshal([]byte(tempBody), &request)
	if err != nil {
		return
	}

	request.Body = []byte(tempFile)

	return
}

func (t transport) PhotoEncode(response models.Photo, ctx *fasthttp.RequestCtx) (err error) {
	panic("implement me")
}
