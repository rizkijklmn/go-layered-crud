package product

// Service berisi logika bisnis
func CreateNewProduct(name string, price int, createdby uint) (*Product, error) {
	product := &Product{Name: name, Price: price, CreatedBy: createdby}
	err := CreateProduct(product)
	return product, err
}

func GetProductsService() ([]Product, error) {
	return GetAllProduct()
}

func GetProductService(id uint) (Product, error) {
	return GetProductByID(id)
}

func UpdateProductService(id uint, name string, price int, createdby uint) (Product, error) {
	product, err := GetProductByID(id)
	if err != nil {
		return product, err
	}

	product.Name = name
	product.Price = price
	product.CreatedBy = createdby

	err = UpdateProduct(&product)
	return product, err
}

func DeleteProductService(id uint) error {
	return DeleteProduct(id)
}
