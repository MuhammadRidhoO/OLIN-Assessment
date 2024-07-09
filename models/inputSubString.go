package models

type InputSubString struct {
	S     string   `json:"s" binding:"required"`
	Words []string `json:"words" binding:"required"`
}