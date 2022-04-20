package server

import (
	"port/pkg/db"
	"port/pkg/util"
	"strconv"
)

func SetupService(handler db.Handler) (*ServiceConfig, error) {
	config := ServiceConfig{}
	port := util.ReadEnvVar("PORT", "50052")
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	config.port = port

	config.url = util.ReadEnvVar("URL", "localhost")
	if handler == nil {
		return nil, ErrNilHandler{}
	}
	config.dbHandler = handler
	return &config, nil
}
