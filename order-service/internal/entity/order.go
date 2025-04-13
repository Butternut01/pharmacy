package entity

type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "pending"
    OrderStatusCompleted OrderStatus = "completed"
    OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
    ProductID string  `json:"product_id" bson:"product_id"`
    Quantity  int     `json:"quantity" bson:"quantity"`
    Price     float64 `json:"price" bson:"price"`
}

type Order struct {
    ID        string      `json:"id" bson:"_id,omitempty"`
    UserID    string      `json:"user_id" bson:"user_id"`
    Items     []OrderItem `json:"items" bson:"items"`
    Total     float64     `json:"total" bson:"total"`
    Status    OrderStatus `json:"status" bson:"status"`
    CreatedAt int64       `json:"created_at" bson:"created_at"`
    UpdatedAt int64       `json:"updated_at" bson:"updated_at"`
}

type OrderFilter struct {
    UserID string
    Status OrderStatus
    Page   int
    Limit  int
}