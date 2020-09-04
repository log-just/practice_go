package data

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// // gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

// User : 유저테이블
type User struct {
	gorm.Model
	Name         string         `gorm:"unique" validate:"required"`
	Email        *string        `validate:"required"`
	Age          uint8          `validate:"required"`
	Birthday     *time.Time     `validate:"required"`
	MemberNumber sql.NullString `validate:"required"`
	ActivedAt    sql.NullTime   `validate:"required"`
}
