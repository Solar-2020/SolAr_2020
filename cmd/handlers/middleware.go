package handlers

type Middleware interface {
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return middleware{}
}
