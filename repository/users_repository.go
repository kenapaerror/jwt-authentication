package repository

import (
	"context"
	"database/sql"
	"yt-users-service/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, user entity.Users)
	FindById(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Users
}
