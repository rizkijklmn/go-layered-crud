package product

import "go-crud/config"

// Repository berinteraksi langsung dengan DB
func CreateProduct(product *Product) error {
	return config.DB.Create(product).Error
}

func GetAllProduct() ([]Product, error) {
	var products []Product
	err := config.DB.Find(&products).Error
	return products, err
}

func GetProductByID(id uint) (Product, error) {
	var product Product
	err := config.DB.First(&product, id).Error
	return product, err
}

func UpdateProduct(product *Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&Product{}, id).Error
}
