package server

import "port/pkg/db"

type ServiceConfig struct {
	port      string
	url       string
	debug     bool
	dbHandler db.Handler
}

type service struct {
	dbHandler db.Handler
}

type ErrNilHandler struct{}
