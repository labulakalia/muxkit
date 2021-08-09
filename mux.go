package muxkit

import (
	"net/http"
)

func defaultNotFoundHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

type Mux struct {
	// middle ware
	// middle will move to route handlerfunc before start
	middle []http.HandlerFunc

	// routre
	// http method / request uri / handerFunc
	route map[string]map[string][]http.HandlerFunc
}

func (m *Mux) addHandle(method string, pattern string, handle ...http.HandlerFunc) {
	if m.route[method] == nil {
		m.route[method] = map[string][]http.HandlerFunc{}
	}
	if m.route[method][pattern] == nil {
		m.route[method][pattern] = []http.HandlerFunc{}
	}
	m.route[method][pattern] = append(m.route[method][pattern], handle...)
}

func (m *Mux) getHandle(method string, pattern string) []http.HandlerFunc {
	_, ok := m.route[method]
	if !ok || len(m.route[method][pattern]) == 0 {
		return []http.HandlerFunc{defaultNotFoundHandlerFunc}
	}
	return m.route[method][pattern]
}

func (s *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range s.getHandle(r.Method, r.URL.Path) {
		handler.ServeHTTP(w, r)
	}
}

func (s *Mux) HandleFunc(httpMethod string, pattern string, handle ...http.HandlerFunc) {
	for _, h := range handle {
		s.addHandle(httpMethod, pattern, append(s.middle, h)...)
	}
}

func (s *Mux) Get(pattern string, handle ...http.HandlerFunc) {
	s.HandleFunc(http.MethodGet, pattern, handle...)
}

func (s *Mux) Post(pattern string, handle ...http.HandlerFunc) {
	s.HandleFunc(http.MethodPost, pattern, handle...)
}

func (s *Mux) Put(pattern string, handle ...http.HandlerFunc) {
	s.HandleFunc(http.MethodPut, pattern, handle...)
}

func (s *Mux) Patch(pattern string, handle ...http.HandlerFunc) {
	s.HandleFunc(http.MethodPatch, pattern, handle...)
}

func (s *Mux) Delete(pattern string, handle ...http.HandlerFunc) {
	s.HandleFunc(http.MethodDelete, pattern, handle...)
}

func (s *Mux) Use(middle ...http.HandlerFunc) {
	s.middle = append(s.middle, middle...)
}

func NewMux() http.Handler {
	return &Mux{}
}
