package exception

import "errors"

var (
	UserConflict = errors.New("User has exist!")
	SignUpFail = errors.New("Fail to create account!")
)