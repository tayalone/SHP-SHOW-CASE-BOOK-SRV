package domains

import (
	"time"
)

// Book is db schema `book`
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Desc      *string   `gorm:"null" json:"desc"`
	Author    string    `gorm:"not null" json:"author"`
	CreatedAt time.Time `gorm:"index;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" json:"updatedAt"`
}
