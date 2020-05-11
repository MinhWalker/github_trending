package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignin(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "walker",
		"email": "minhnh@gmail.com",
	})
}

func HandleSignup(c echo.Context) error {
	type User struct {
		Email string	`json:"email"`
		Fullname string	`json:"fullname"`
		Age int			`json:"age"`
	}

	user := User{
		Email: "minhnh@gmail.com",
		Fullname: "Minh Walker",
		Age: 20,
	}
	return c.JSON(http.StatusOK, user)
}
