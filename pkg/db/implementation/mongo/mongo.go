package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	i "port/pkg/db/implementation"
	"port/pkg/grpc/pb"
)

func New(dbConfig map[i.DbParam]string) (*Database, error) {
	log.Info("Mongo Database")
	uri := "mongodb://" + dbConfig[i.DbUser] + ":" + dbConfig[i.DbPassword] + "@" + dbConfig[i.DbUrl]

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

	return &Database{dbName: dbConfig[i.DbName], collection: dbConfig[i.DbCollection], client: client}, nil
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
