package security

import (
	"backend-github-trending/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRET_KEY = "dgdsgsdgf3fdsf3dwdw"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //gen token, insert algorithm
	tokenString, err := token.SignedString([]byte(SECRET_KEY))      //ma hoa token
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
