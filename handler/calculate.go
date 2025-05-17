package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SumHandler godoc
// @Summary Вычисление суммы
// @Accept json
// @Produce json
// @Param numbers body map[string]int true "Объект с числами"
// @Success 200 {object} map[string]int
// @Router /calculate [post]
func SumHandler(c *gin.Context) {
	var numbers map[string]int
	if err := c.ShouldBindJSON(&numbers); err != nil {
		c.JSON(400,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	c.JSON(http.StatusOK,
		gin.H{
			"result": sum,
		})
}
