package handler

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// QueryHandler godoc
// @Summary Выполнить SQL по CSV
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV-файл"
// @Param sql formData string true "SQL-запрос"
// @Success 200 {object} map[string]any
// @Router /query [post]
func QueryHandler(c *gin.Context) {
	query := c.PostForm("sql")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query is empty",
		})
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()

	csvData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query = normalizeQuery(query)
	cmd := exec.Command("q", "-H", "-d", ",", query)
	cmd.Stdin = bytes.NewReader(csvData)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"styderr": stderr.String(),
		})
		return
	}

	reader := csv.NewReader(strings.NewReader(string(output)))
	rows, err := reader.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rows == nil {
		rows = [][]string{}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": rows,
	})
}

func normalizeQuery(query string) string {
	lower := strings.ToLower(query)

	if !strings.Contains(lower, "from") {
		if idx := strings.Index(lower, "where"); idx != -1 {
			return query[:idx] + "FROM - " + query[idx:]
		} else if idx := strings.Index(lower, "limit"); idx != -1 {
			return query[:idx] + "FROM - " + query[idx:]
		} else if idx := strings.Index(lower, "order by"); idx != -1 {
			return query[:idx] + "FROM - " + query[idx:]
		}

		return query + " FROM -"
	}

	return query
}
