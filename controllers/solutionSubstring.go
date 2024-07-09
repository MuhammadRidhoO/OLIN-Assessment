package controllers

import (
	"fmt"
	"net/http"
	"task-company/models"

	"github.com/gin-gonic/gin"
)

func FindSubstring(c *gin.Context) {
    var input models.InputSubString
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := findSubstring(input.S, input.Words)
    c.JSON(http.StatusOK, gin.H{"indices": result})
}

func findSubstring(s string, words []string) []int {
    if len(words) == 0 || len(s) == 0 {
        return []int{}
    }

    wordLen := len(words[0])
    wordCount := len(words)
    concatLen := wordLen * wordCount
    wordMap := make(map[string]int)

    for _, word := range words {
        wordMap[word]++
    }

    result := []int{}
    fmt.Printf("Word Map: %v\n", wordMap)

    for i := 0; i <= len(s)-concatLen; i++ {
        seen := make(map[string]int)
        j := 0
        for ; j < wordCount; j++ {
            start := i + j*wordLen
            end := start + wordLen
            if end > len(s) {
                break
            }
            word := s[start:end]
            fmt.Printf("Checking word: %s at position %d\n", word, i)
            if count, ok := wordMap[word]; ok {
                seen[word]++
                if seen[word] > count {
                    break
                }
            } else {
                break
            }
        }
        if j == wordCount {
            result = append(result, i)
        }
    }
    fmt.Printf("Result indices: %v\n", result)
    return result
}
