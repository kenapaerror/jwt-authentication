package service

import (
	"context"
	"yt-users-service/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse
	Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) web.UsersResponse
	FindAll(ctx context.Context) []web.UsersResponse
	Auth(ctx context.Context, request web.UserAuthRequest) web.TokenResponse
	CreateWithRefreshToken(ctx context.Context, refreshToken string) web.TokenResponse
}
