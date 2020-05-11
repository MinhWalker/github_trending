package main

import (
	"backend-github-trending/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignin)
	e.GET("/user/sign-up", handler.HandleSignup)

	e.Logger.Fatal(e.Start(":3000"))
}



