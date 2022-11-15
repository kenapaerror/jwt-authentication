package repository

import (
	"context"
	"database/sql"
	"errors"
	"yt-users-service/helper"
	"yt-users-service/model/entity"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	helper.PanicError(err)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "UPDATE users SET first_name=?, last_name=?, updated_at=? WHERE id=?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.Id,
	)
	helper.PanicError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user entity.Users) {
	SQL := "DELETE FROM users WHERE id=?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
	)
	helper.PanicError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error) {
	SQL := "SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		userId,
	)
	helper.PanicError(err)
	defer rows.Close()

	user := entity.Users{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error) {
	SQL := "SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE email = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		email,
	)
	helper.PanicError(err)
	defer rows.Close()

	user := entity.Users{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Users {
	SQL := "SELECT id, first_name, last_name, email, created_at, updated_at FROM users"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicError(err)
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		user := entity.Users{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicError(err)

		users = append(users, user)
	}

	return users
}
