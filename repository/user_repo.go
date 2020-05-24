package repository

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"context"
)

type UserRepo interface {
	SaveUser(ctx context.Context, user model.User) (model.User, error)
	CheckLogin(ctx context.Context, loginReq req.RepSignIn) (model.User, error)
}
