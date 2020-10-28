package handlers

import (
	httputils "github.com/Solar-2020/GoUtils/http"
	postsHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/posts"
	uploadHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/upload"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(posts postsHandler.Handler, upload uploadHandler.Handler, middleware httputils.Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.PanicHandler = httputils.PanicHandler
	clientside := httputils.ClientsideChain(middleware)

	router.Handle("GET", "/health", middleware.Log(httputils.HealthCheckHandler))

	router.Handle("POST", "/api/posts/post", clientside(posts.Create))
	router.Handle("GET", "/api/posts/posts", clientside(posts.GetList))

	router.Handle("POST", "/api/upload/photo", clientside(upload.Photo))
	router.Handle("POST", "/api/upload/file", clientside(upload.File))

	return router
}
