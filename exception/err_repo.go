package exception

import "errors"

var (
	//repo
	RepoNotUpdate = errors.New("Fail to update infomation")
	RepoNotFound  = errors.New("Repo not exist")
	RepoConlict   = errors.New("Repo has exist")
	RepoInsertFail = errors.New("Fail to add Repo")
)


