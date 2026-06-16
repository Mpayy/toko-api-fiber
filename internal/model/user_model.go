package model

import "time"

type UserResponse struct {
	ID        int64      `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `json:"token" validate:"required"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"-" validate:"required,min=1"`
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type LogoutUserRequest struct {
	ID int64 `json:"id" validate:"required,min=1"`
}

type GetUserRequest struct {
	ID int64 `json:"id" validate:"required,min=1"`
}
