package handlers

import (
	postsHandler "github.com/Solar-2020/SolAr_2020/cmd/handlers/posts"
	uploadHandler "github.com/Solar-2020/SolAr_2020/cmd/handlers/upload"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(posts postsHandler.Handler, upload uploadHandler.Handler, middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	//router.Handle("GET", "/health", check)

	router.Handle("POST", "/api/posts/post", middleware.CORS(posts.Create))
	router.Handle("GET", "/api/posts/posts", middleware.CORS(posts.GetList))

	router.Handle("POST", "/api/upload/photo", middleware.CORS(upload.Photo))
	router.Handle("POST", "/api/upload/file", middleware.CORS(upload.File))

	return router
}
