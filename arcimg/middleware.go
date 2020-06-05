package arcimg

import "net/http"

type Middleware struct {
	use    http.HandlerFunc
	Middle []middleware
}

type Http struct {
	i   int
	w   http.ResponseWriter
	req *http.Request
	*Middleware
}

type middleware func(m *Http)

func NewMiddleware(h http.HandlerFunc) Middleware {
	list := make([]middleware, 0)
	return Middleware{use: h, Middle: list}
}

func (m *Middleware) Add(mm middleware) {
	m.Middle = append(m.Middle, mm)
}

func (Midd *Middleware) Use(w http.ResponseWriter, req *http.Request) {
	var m = Http{
		w:          w,
		req:        req,
		Middleware: Midd,
	}
	m.Next()
}

func (m *Http) Next() {
	m.i++
	if m.i-1 < len(m.Middle) {
		m.Middle[m.i-1](m)
	} else {
		m.use(m.w, m.req)
	}
}
