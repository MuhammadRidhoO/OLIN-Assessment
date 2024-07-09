package models

type InputThreeSum struct {
	Nums   []int `json:"nums" binding:"required"`
}