package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"main/internal/main/models"

	"main/pkg/jwt"
)

func Register(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	_, err := models.UserCreate(requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"success": true,
			"message": "Your account is successfully created!",
		},
	})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	if err := c.Bind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	exist, user, _ := models.UserAuthenticate(requestBody.Email, requestBody.Password)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Wrong username or password",
		})
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"success":      true,
			"access_token": token,
			"user":         user,
		},
	})
}

func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "No user found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func GetCreditLimit(c *gin.Context) {
	user, _ := c.Get("user")

	limits, err := models.GetCreditLimit(user.(*models.User).ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    limits,
	})
}

func GetAllContracts(c *gin.Context) {
	user, _ := c.Get("user")

	contracts, err := models.GetAllContracts(user.(*models.User).ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    contracts,
	})
}

func GetContractByID(c *gin.Context) {
	user, _ := c.Get("user")

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	contract, err := models.GetContractByID(id, user.(*models.User).ID)
	if contract == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    contract,
	})
}
