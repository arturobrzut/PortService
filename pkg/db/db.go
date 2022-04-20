package db

import (
	"port/pkg/db/implementation/memory"
	"port/pkg/db/implementation/mongo"
	"port/pkg/util"
	"strings"
)

func NewDbHandler() (Handler, error) {
	switch strings.ToUpper(util.ReadEnvVar("DB_TYPE", "MEMORY")) {
	case "MEMORY":
		return memory.New()
	case "MONGO":
		return mongo.New()
	}
	return nil, ErrUnknownDbType{}
}

func (e ErrUnknownDbType) Error() string {
	return "Unknown database type"
}
