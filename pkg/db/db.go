package db

import (
	implement "port/pkg/db/implementation"
	"port/pkg/db/implementation/memory"
	"port/pkg/db/implementation/mongo"
	"port/pkg/util"
	"strings"
)

func NewDbHandler() (Handler, error) {
	dbConfig := setDbConfig()
	switch strings.ToUpper(dbConfig[implement.DbType]) {
	case "MEMORY":
		return memory.New(dbConfig)
	case "MONGO":
		return mongo.New(dbConfig)
	}
	return nil, ErrUnknownDbType{}
}

func setDbConfig() map[implement.DbParam]string {
	dbConfig := map[implement.DbParam]string{}
	dbConfig[implement.DbType] = util.ReadEnvVar("DB_TYPE", "MEMORY")
	dbConfig[implement.DbUrl] = util.ReadEnvVar("DB_URL", "mongo:27017")
	dbConfig[implement.DbName] = util.ReadEnvVar("DB_NAME", "port_db")
	dbConfig[implement.DbUser] = util.ReadEnvVar("DB_USER", "root")
	dbConfig[implement.DbPassword] = util.ReadEnvVar("DB_PASSWORD", "example")
	dbConfig[implement.DbCollection] = util.ReadEnvVar("DB_COLLECTION", "collection")
	return dbConfig
}

func (e ErrUnknownDbType) Error() string {
	return "Unknown database type"
}
