package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller menangani request dari client
func CreateUserHandler(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := CreateNewUser(input.Name, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsersHandler(c *gin.Context) {
	users, err := GetUsersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := GetUserService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := UpdateUserService(uint(id), input.Name, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := DeleteUserService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
