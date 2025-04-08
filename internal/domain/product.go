package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description" json:"description"`
    Price       float64            `bson:"price" json:"price"`
    Category    string             `bson:"category" json:"category"`
    Stock       int                `bson:"stock" json:"stock"`
}
