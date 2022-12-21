package mongodb

import (
	"context"
	"time"

	"github.com/DLzer/go-echo-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

type MongoDb struct {
	MongoClient *mongo.Client
}

func NewMongoDB(c *config.Config) (*MongoDb, error) {
	// Mongo config
	opt := options.Client().ApplyURI(c.MongoDB.MongoURI).
		SetConnectTimeout(connectTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	// Establish connection
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}

	// Start context with timer
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	// Ping
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoDb{MongoClient: client}, nil
}

func (m *MongoDb) Close() error {
	return m.MongoClient.Disconnect(context.Background())
}
