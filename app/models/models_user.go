package models

import (
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type GlobalUser struct {
	CustomGormModel
	Avatar          string    `json:"avatar" gorm:"type: text"`
	Fullname        string    `json:"fullname" gorm:"type: varchar(255)"`
	Email           string    `json:"email" gorm:"type: varchar(255)"`
	Password        string    `json:"-" gorm:"type: varchar(255)"`
	EmailVerifiedAt null.Time `json:"email_verified_at" gorm:"type: timestamp"`
	Gender          string    `json:"gender" gorm:"column:gender"`
}

func (GlobalUser) TableName() string {
	return "global_users"
}

type CustomGormModel struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	EncodedID string          `json:"encoded_id" gorm:"-"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type UserServerResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    GlobalUser `json:"data"`
}
