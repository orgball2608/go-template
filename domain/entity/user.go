package entity

import "time"

type User struct {
	Id                   int       `json:"id"`
	Email                string    `json:"email"`
	Type                 string    `json:"type"`
	Name                 string    `json:"fullName"`
	IsMiddleManager      bool      `json:"isMiddleManager"`
	IsProxy              bool      `json:"isPOroxy"`
	IsHelpDeskUser       bool      `json:"isHelpDeskUser"`
	Password             string    `json:"password"`
	Status               string    `json:"status"`
	DepartmentName       string    `json:"departmentName"`
	LastModifiedBy       string    `json:"lastModifiedBy"`
	LastLoginAt          time.Time `json:"lastLoginAt"`
	LastChangePasswordAt time.Time `json:"lastChangePasswordAt"`
	OwnerId              string    `json:"ownerId"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
	DeletedAt            time.Time `json:"deletedAt"`
}
