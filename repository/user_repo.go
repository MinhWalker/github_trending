package repository

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"context"
)

type UserRepo interface {
	//user
	SaveUser(ctx context.Context, user model.User) (model.User, error)
	CheckLogin(ctx context.Context, loginReq req.RepSignIn) (model.User, error)

	//profile
	SelectUserById(ctx context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
}
