package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"usermovieratingapi/config"
)

type DB struct{
	connection *mongo.Database
	connectionError error
}

var runOnce sync.Once

func (db *DB) SetupConnection(conf *config.Configuration) (*mongo.Database, error) {
	// Perform connection creation operation only once.
	runOnce.Do(func() {
		// Set client options
		credential := options.Credential{
			Username: conf.User,
			Password: conf.Password,
		}
		clientOptions := options.Client().ApplyURI(conf.CONNECTIONSTRING).SetAuth(credential)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			db.connectionError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			db.connectionError = err
		}
		db.connection = client.Database(conf.DATABASE)
	})

	return db.connection, db.connectionError
}

func (db *DB) GetCollection(name string) (*mongo.Collection, error) {
	if db.connectionError != nil {
		return nil, db.connectionError
	}
	return db.connection.Collection(name), db.connectionError
}
