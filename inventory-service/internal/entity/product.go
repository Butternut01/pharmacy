package entity

type Product struct {
    ID          string  `json:"id" bson:"_id,omitempty"`
    Name        string  `json:"name" bson:"name"`
    Description string  `json:"description" bson:"description"`
    Price       float64 `json:"price" bson:"price"`
    Stock       int     `json:"stock" bson:"stock"`
    Category    string  `json:"category" bson:"category"`
}

type ProductFilter struct {
    Name     string
    Category string
    MinPrice float64
    MaxPrice float64
    Page     int
    Limit    int
}