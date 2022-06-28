package model

type Document struct {
	ID          uint   `json:"id" binding:"required" gorm:"AUTO_INCREMENT;primary_key;index"`
	Name        string `json:"name" binding:"required" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
}
