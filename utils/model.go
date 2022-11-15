package utils

import (
	"yt-users-service/model/entity"
	"yt-users-service/model/web"
)

func UserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
