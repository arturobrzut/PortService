package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	client     *mongo.Client
	dbName     string
	collection string
}
