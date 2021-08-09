package http

import "time"

type httpOption struct {
	Addr string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	
}

// func NewServer()
