package models

type InputTwoSum struct {
	Nums   []int `json:"nums" binding:"required"`
	Target int   `json:"target" binding:"required"`
}