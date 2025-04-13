package config

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
    MongoDBURI      string
    MongoDBName     string
    ServerPort      string
}

func NewConfig() *Config {
    return &Config{
        MongoDBURI:      "mongodb://localhost:27017",
        MongoDBName:     "inventory_db",
        ServerPort:      "8080",
    }
}

func ConnectMongoDB(uri string) (*mongo.Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }

    if err = client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
    }

    return client.Database("inventory_db"), nil
}