package repository

import (
    "context"
    "inventory-service/internal/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"  // Добавь этот импорт
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

// Create новый продукт
func (r *ProductRepository) Create(product domain.Product) (*domain.Product, error) {
	result, err := r.collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return &product, nil
}

// FindByID находит продукт по ID
func (r *ProductRepository) FindByID(id primitive.ObjectID) (*domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Update обновляет продукт
func (r *ProductRepository) Update(id primitive.ObjectID, product domain.Product) (*domain.Product, error) {
	update := bson.M{
		"$set": product,
	}

	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	product.ID = id
	return &product, nil
}

// Delete удаляет продукт по ID
func (r *ProductRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// FindAll находит все продукты с пагинацией
func (r *ProductRepository) FindAll(page, pageSize int) ([]domain.Product, error) {
	var products []domain.Product

	options := options.Find()
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
