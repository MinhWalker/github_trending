package repo_impl

import (
	"backend-github-trending/db"
	"backend-github-trending/log"
	"backend-github-trending/exception"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, full_name, email, password, role, created_at, updated_at)
		VALUES (:user_id, :full_name, :email, :password, :role, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(ctx, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, exception.UserConflict
			}
		}
		return user, exception.SignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) CheckLogin(ctx context.Context, loginReq req.RepSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(ctx, &user,"SELECT * FROM users WHERE email=$1", loginReq.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, exception.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}