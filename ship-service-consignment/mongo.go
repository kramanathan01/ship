package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create - Adds a new Consignment with one or more Containers
func (repository *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll - Returns all Consignments
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repository.collection.Find(ctx, bson.D{}, nil)
	defer cur.Close(ctx)
	if err != nil {
		log.Printf("Error in finding - Context: %v, Error: %v", ctx, err)
		return nil, err
	}

	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}

// CreateClient -
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err := conn.Ping(ctx, nil); err != nil {
		if retry >= 3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}
