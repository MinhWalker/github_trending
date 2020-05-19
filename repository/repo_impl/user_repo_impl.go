package repo_impl

import (
	"backend-github-trending/db"
	"backend-github-trending/log"
	"backend-github-trending/exception"
	"backend-github-trending/model"
	"backend-github-trending/repository"
	"context"
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

func (u UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
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