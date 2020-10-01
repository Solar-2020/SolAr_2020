package handlers

import (
	postsHandler "github.com/BarniBl/SolAr_2020/cmd/handlers/posts"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(posts postsHandler.Handler, middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	//router.Handle("GET", "/health", check)

	router.Handle("POST", "/posts/post", posts.Create)
	router.Handle("GET", "/posts/posts", posts.GetList)

	router.Handle("GET", "/upload", posts.GetList)

	return router
}
