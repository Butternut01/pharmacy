package repository

import (
    "context"
    "time"

    "github.com/Butternut01/inventory-service/internal/entity"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
    Create(product *entity.Product) error
    FindByID(id string) (*entity.Product, error)
    Update(product *entity.Product) error
    Delete(id string) error
    FindAll(filter entity.ProductFilter) ([]entity.Product, error)
}

type productRepository struct {
    collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
    return &productRepository{
        collection: db.Collection("products"),
    }
}

func (r *productRepository) Create(product *entity.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.InsertOne(ctx, product)
    return err
}

func (r *productRepository) FindByID(id string) (*entity.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var product entity.Product
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
    if err != nil {
        return nil, err
    }

    return &product, nil
}

func (r *productRepository) Update(product *entity.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(product.ID)
    if err != nil {
        return err
    }

    update := bson.M{
        "$set": bson.M{
            "name":        product.Name,
            "description": product.Description,
            "price":       product.Price,
            "stock":       product.Stock,
            "category":    product.Category,
        },
    }

    _, err = r.collection.UpdateByID(ctx, objectID, update)
    return err
}

func (r *productRepository) Delete(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    return err
}

func (r *productRepository) FindAll(filter entity.ProductFilter) ([]entity.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    query := bson.M{}
    if filter.Name != "" {
        query["name"] = bson.M{"$regex": filter.Name, "$options": "i"}
    }
    if filter.Category != "" {
        query["category"] = filter.Category
    }
    if filter.MinPrice > 0 || filter.MaxPrice > 0 {
        priceQuery := bson.M{}
        if filter.MinPrice > 0 {
            priceQuery["$gte"] = filter.MinPrice
        }
        if filter.MaxPrice > 0 {
            priceQuery["$lte"] = filter.MaxPrice
        }
        query["price"] = priceQuery
    }

    opts := options.Find()
    if filter.Limit > 0 {
        opts.SetLimit(int64(filter.Limit))
        opts.SetSkip(int64((filter.Page - 1) * filter.Limit))
    }

    cursor, err := r.collection.Find(ctx, query, opts)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []entity.Product
    if err = cursor.All(ctx, &products); err != nil {
        return nil, err
    }

    return products, nil
}