package config

type Config struct {
    ServerPort          string
    InventoryServiceURL string
    OrderServiceURL     string
}

func NewConfig() *Config {
    return &Config{
        ServerPort:          "8000",
        InventoryServiceURL: "http://localhost:8080",
        OrderServiceURL:     "http://localhost:8081",
    }
}