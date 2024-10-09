package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) create(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"name": name,
			"id":   id,
		},
	})
	DB.Create(&IdCard{Name: name, Id: id})
}
func (app *application) view(c *gin.Context) {
	id := c.Query("id")

	var idCard IdCard
	// 尝试从数据库中查询记录
	result := DB.First(&idCard, "id = ?", id)
	// 检查是否有记录被找到
	if result.Error != nil {
		// 如果查询出错，返回错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		// 如果没有找到记录，返回404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": "IdCard not found"})
		return
	}
	// 如果查询成功，返回查询结果
	c.JSON(http.StatusOK, idCard)
}
