package usecase

import (
    "github.com/Butternut01/inventory-service/internal/entity"
    "github.com/Butternut01/inventory-service/internal/repository"
)

type ProductUseCase interface {
    CreateProduct(product *entity.Product) error
    GetProduct(id string) (*entity.Product, error)
    UpdateProduct(product *entity.Product) error
    DeleteProduct(id string) error
    ListProducts(filter entity.ProductFilter) ([]entity.Product, error)
}

type productUseCase struct {
    productRepo repository.ProductRepository
}

func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
    return &productUseCase{
        productRepo: productRepo,
    }
}

func (uc *productUseCase) CreateProduct(product *entity.Product) error {
    return uc.productRepo.Create(product)
}

func (uc *productUseCase) GetProduct(id string) (*entity.Product, error) {
    return uc.productRepo.FindByID(id)
}

func (uc *productUseCase) UpdateProduct(product *entity.Product) error {
    return uc.productRepo.Update(product)
}

func (uc *productUseCase) DeleteProduct(id string) error {
    return uc.productRepo.Delete(id)
}

func (uc *productUseCase) ListProducts(filter entity.ProductFilter) ([]entity.Product, error) {
    return uc.productRepo.FindAll(filter)
}