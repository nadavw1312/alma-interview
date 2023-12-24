package dal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BaseFindById[T any](ctx context.Context, collection *mongo.Collection, id string) (T, error) {
	var doc T
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return doc, err

	}

	query := bson.D{{Key: "_id", Value: oid}}
	err = collection.FindOne(ctx, query).Decode(&doc)

	return doc, err
}

func BaseInsert(ctx context.Context, collection *mongo.Collection, doc any) (string, error) {
	insertionResult, err := collection.InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	return insertionResult.InsertedID.(primitive.ObjectID).Hex(), err
}

func BaseDeleteById(ctx context.Context, collection *mongo.Collection, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	if _, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: oid}}); err != nil {
		return err
	}

	return nil
}

func BaseFindOneByQuery[T any](ctx context.Context, collection *mongo.Collection, query interface{}) (T, error) {
	var doc T
	err := collection.FindOne(ctx, query).Decode(&doc)
	return doc, err
}
func BaseFindByQuery[T any](ctx context.Context, collection *mongo.Collection, query interface{}) ([]T, error) {
	var doc []T
	res, err := collection.Find(ctx, query)
	res.All(ctx, &doc)
	return doc, err
}
