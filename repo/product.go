package repo

import (
	"github.com/cladamos/clamarket-be/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	result := r.db.Create(product)
	return result.Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	result := r.db.First(&product, "id = ?", id)
	return &product, result.Error
}
