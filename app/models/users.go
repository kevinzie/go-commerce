package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	//gorm.Model
	ID              uint           `db:"id" json:"-" validate:"" gorm:"primary_key"`    // column name is `id`
	Uuid            uuid.UUID      `db:"uuid" json:"user_id"  validate:"required,uuid"` // column name is `uuid`
	FirstName       string         `db:"first_name" json:"first_name"  xml:"first_name" form:"first_name" validate:"required, lte=255"`
	LastName        string         `db:"last_name" json:"last_name"  xml:"last_name" form:"last_name" validate:"lte=255"`
	Email           string         `db:"email" json:"email"  xml:"email" form:"email" validate:"lte=255"`
	EmailVerifiedAt *time.Time     `db:"email_verified_at" json:"email_verified_at" validate:"lte=255"`
	Password        string         `db:"password" json:"-"  xml:"password" form:"password" gorm:"type:varchar(200)" validate:"required"`
	RememberToken   *string        `db:"remember_token" json:"remember_token" validate:""`
	Activated       uint           `db:"activated" json:"activated" validate:""`
	Status          string         `db:"status" json:"status" validate:"" gorm:"default:inactive"`
	SignupIpAddress *string        `db:"signup_ip_address" json:"signup_ip_address" validate:""`
	Profiles        Profiles       `json:"profile" gorm:"foreignKey:UserId"`
	CreatedAt       time.Time      `db:"created_at" json:"created_at" validate:""`
	UpdatedAt       *time.Time     `db:"updated_at" json:"updated_at" validate:""`
	DeletedAt       gorm.DeletedAt `db:"deleted_at" json:"-" gorm:"index"`
}

//func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//	u.Password = fmt.Sprintf("%x", bytes)
//	return nil
//}

//
//func (u *Users) BeforeSave(tx *gorm.DB) (err error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
//	u.Password = fmt.Sprintf("%x", bytes)
//	return nil
//}
