package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

var validationMessages = map[string]string{
	"required": "The :field is required",
	"uri":      "The :field is not a valid URL",
}

func ValidateRequest(c *gin.Context, obj any) bool {
	if err := c.ShouldBind(&obj); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			rule := fieldErr.Tag()

			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": strings.ReplaceAll(validationMessages[rule], ":field", fieldErr.Field()),
			})
			return false
		}
	}

	return true
}
