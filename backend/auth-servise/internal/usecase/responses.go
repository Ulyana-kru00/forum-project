package usecase

import "github.com/Ulyana-kru00/forum-project/internal/entity"

type RegisterResponse struct {
	UserID int64
}

type LoginResponse struct {
	Token    string
	Username string
}

type ValidateTokenResponse struct {
	UserID int64
	Role   string
	Valid  bool
}

type GetUserResponse struct {
	User *entity.User
}
