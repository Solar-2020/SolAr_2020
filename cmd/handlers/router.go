package handlers

import (
	httputils "github.com/Solar-2020/GoUtils/http"
	postsHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/posts"
	uploadHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/upload"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(posts postsHandler.Handler, upload uploadHandler.Handler, middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.PanicHandler = httputils.PanicHandler

	router.Handle("GET", "/health", middleware.Log(httputils.HealthCheckHandler))

	router.Handle("POST", "/api/posts/post", middleware.Log(middleware.ExternalAuth(posts.Create)))
	//router.Handle("POST", "/api/posts/post", posts.Create)
	router.Handle("GET", "/api/posts/posts", middleware.Log(middleware.ExternalAuth(posts.GetList)))

	router.Handle("POST", "/api/upload/photo", middleware.Log(middleware.ExternalAuth(upload.Photo)))
	router.Handle("POST", "/api/upload/file", middleware.Log(middleware.ExternalAuth(upload.File)))

	return router
}
