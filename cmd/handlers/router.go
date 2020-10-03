package handlers

import (
	postsHandler "github.com/BarniBl/SolAr_2020/cmd/handlers/posts"
	uploadHandler "github.com/BarniBl/SolAr_2020/cmd/handlers/upload"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(posts postsHandler.Handler, upload uploadHandler.Handler,  middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	//router.Handle("GET", "/health", check)

	router.Handle("POST", "/api/posts/post", posts.Create)
	router.Handle("GET", "/api/posts/posts", posts.GetList)

	router.Handle("POST", "/api/upload/photo", upload.Photo)
	router.Handle("POST", "/api/upload/file", upload.File)

	return router
}
