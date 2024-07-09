package controllers

import (
	"net/http"
	"task-company/models"

	"github.com/gin-gonic/gin"
)

func Towsum(c *gin.Context){
	var input models.InputTwoSum
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	result := findTwoSum(input.Nums, input.Target)
    if result == nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "No solution found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"indices": result})
}

func findTwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, ok := m[complement]; ok {
            return []int{idx, i}
        }
        m[num] = i
    }
    return nil
}
