package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"port/pkg/grpc/pb"
	"port/pkg/util"
)

func New() (*Database, error) {
	log.Info("Mongo Database")
	dbConfig := setDbConfig()
	uri := "mongodb://" + dbConfig[DbUser] + ":" + dbConfig[DbPassword] + "@" + dbConfig[DbUrl]

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Error("Error during connection to mongo. ", err)
		return nil, err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Error("Cannot ping to the mongo instance. ", err)
		return nil, err
	}
	log.Info("Mongo Ping successfully")

	return &Database{dbName: dbConfig[DbName], collection: dbConfig[DbCollection], client: client}, nil
}

func (db *Database) Close() {
	log.Print("Close Mongo")
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Error("Error during mongo disconnect")
	}
}

func (db *Database) Create(port *pb.Port) error {
	panic("not implemented yet")
}

func (db *Database) Get(id string) (*pb.Port, error) {
	panic("not implemented yet")
}

func (db *Database) Delete(id string) error {
	panic("not implemented yet")
}

func (db *Database) Update(port *pb.Port) error {
	panic("not implemented yet")
}

func (db *Database) CreateOrUpdate(port *pb.Port) error {
	panic("not implemented yet")
}

func setDbConfig() map[int]string {
	dbConfig := map[int]string{}
	dbConfig[DbUrl] = util.ReadEnvVar("DB_URL", "mongo:27017")
	dbConfig[DbName] = util.ReadEnvVar("DB_NAME", "port_db")
	dbConfig[DbUser] = util.ReadEnvVar("DB_USER", "root")
	dbConfig[DbPassword] = util.ReadEnvVar("DB_PASSWORD", "example")
	dbConfig[DbCollection] = util.ReadEnvVar("DB_COLLECTION", "collection")
	return dbConfig
}
