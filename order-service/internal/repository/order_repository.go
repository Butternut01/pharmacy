package repository

import (
    "context"
    "time"

    "github.com/Butternut01/order-service/internal/entity"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository interface {
    Create(order *entity.Order) error
    FindByID(id string) (*entity.Order, error)
    UpdateStatus(id string, status entity.OrderStatus) error
    FindAll(filter entity.OrderFilter) ([]entity.Order, error)
}

type orderRepository struct {
    collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) OrderRepository {
    return &orderRepository{
        collection: db.Collection("orders"),
    }
}

func (r *orderRepository) Create(order *entity.Order) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    order.CreatedAt = time.Now().Unix()
    order.UpdatedAt = order.CreatedAt

    _, err := r.collection.InsertOne(ctx, order)
    return err
}

func (r *orderRepository) FindByID(id string) (*entity.Order, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var order entity.Order
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
    if err != nil {
        return nil, err
    }

    return &order, nil
}

func (r *orderRepository) UpdateStatus(id string, status entity.OrderStatus) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    update := bson.M{
        "$set": bson.M{
            "status":     status,
            "updated_at": time.Now().Unix(),
        },
    }

    _, err = r.collection.UpdateByID(ctx, objectID, update)
    return err
}

func (r *orderRepository) FindAll(filter entity.OrderFilter) ([]entity.Order, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    query := bson.M{}
    if filter.UserID != "" {
        query["user_id"] = filter.UserID
    }
    if filter.Status != "" {
        query["status"] = filter.Status
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

    var orders []entity.Order
    if err = cursor.All(ctx, &orders); err != nil {
        return nil, err
    }

    return orders, nil
}