package mongo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/teltech/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

const CONNECTED = "Successfully connected to database: %s"

type Datastore struct {
	Db      *mongo.Database
	Session *mongo.Client
	logger  *logger.Log
}

func New(connectionString string, timeout int, databaseName string, log *logger.Log) *Datastore {

	var mongoDataStore *Datastore
	db, session := connect(connectionString, timeout, databaseName, log)
	if db != nil && session != nil {

		mongoDataStore = new(Datastore)
		mongoDataStore.Db = db
		mongoDataStore.logger = log
		mongoDataStore.Session = session
		return mongoDataStore
	}


	return nil
}

func connect(connectionString string, timeout int, databaseName string, log *logger.Log) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(connectionString, timeout, databaseName, log)
	})

	return db, session
}

func connectToMongo(connectionString string, timeout int, databaseName string, log *logger.Log) (a *mongo.Database, b *mongo.Client) {

	var err error

	clientOptions := options.Client().SetSocketTimeout(time.Duration(timeout) * time.Second).ApplyURI(connectionString)
	session, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("error connecting to mongo: %s.", connectionString)))
	}

	err = session.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("error pinging mongo: %s.", connectionString )))
	}

	log.Info(CONNECTED);

	var DB = session.Database(databaseName)

	return DB, session
}
