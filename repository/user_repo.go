package repository

import (
	"backend-github-trending/model"
	"context"
)

type UserRepo interface {
	SaveUser(ctx context.Context, user model.User) (model.User, error)
}