package repository

import (
	"context"
	"inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{collection: db.Collection("categories")}
}

func (r *CategoryRepository) Create(category domain.Category) (*domain.Category, error) {
	res, err := r.collection.InsertOne(context.Background(), category)
	if err != nil {
		return nil, err
	}
	category.ID = res.InsertedID.(primitive.ObjectID)
	return &category, nil
}

func (r *CategoryRepository) FindByID(id primitive.ObjectID) (*domain.Category, error) {
	var category domain.Category
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindAll() ([]domain.Category, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var categories []domain.Category
	for cursor.Next(context.Background()) {
		var category domain.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}
	return categories, nil
}
