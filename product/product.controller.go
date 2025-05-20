package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller menangani request dari client
func CreateProductHandler(c *gin.Context) {
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := CreateNewProduct(input.Name, input.Price, input.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal create product"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductsHandler(c *gin.Context) {
	products, err := GetProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := GetProductService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := UpdateProductService(uint(id), input.Name, input.Price, input.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := DeleteProductService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product berhasil dihapus"})
}
