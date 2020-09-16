package handler

import (
	"backend-github-trending/log"
	"backend-github-trending/model"
	req "backend-github-trending/model/req"
	"backend-github-trending/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type RepoHandler struct {
	GithubRepo repository.GithubRepo
}

// RepoTrending godoc
// @Summary Get all repo Trending on
// @Tags -service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Router /github/trending [GET]
// @return Repo trending

func (r RepoHandler) RepoTrending(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), claims.UserId, 25)
	for i, repo := range repos {
		repos[i].Contributors = strings.Split(repo.BuildBy, ",")
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       repos,
	})
}

// RepoTrending godoc
// @Summary Get all repo Trending on
// @Tags -service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Router /github/trending [GET]
// @return Repo trending
func (r RepoHandler) SelectBookmarks(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	repos, _ := r.GithubRepo.SelectAllBookmarks(c.Request().Context(), claims.UserId)

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success!",
		Data: repos,
	})
}

// InSertBookmark godoc
// @Summary Insert Repo to bookmark table
// @Tags -service
// @Accept  json
// @Produce  json
// @Param data body req.ReqBookmark true "GithubRepo"
// @Success 200 {object} model.Response
// @Success 403 {object} model.Response
// @Success 400 {object} model.Response
// @Success 500 {object} model.Response
// @Router /bookmark/add [POST]

func (r RepoHandler) InsertBookmark(c echo.Context) error {
	req := req.ReqBookmark{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	bId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = r.GithubRepo.InsertBookmark(
		c.Request().Context(),
		bId.String(),
		req.RepoName,
		claims.UserId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Bookmark thành công",
		Data:       nil,
	})
}

// DelBookmark godoc
// @Summary Delete Repo to bookmark table
// @Tags -service
// @Accept  json
// @Produce  json
// @Param data body req.ReqBookmark true "GithubRepo"
// @Success 200 {object} model.Response
// @Success 400 {object} model.Response
// @Success 500 {object} model.Response
// @Router /bookmark/delete [DELETE]

func (r RepoHandler) DelBookmark(c echo.Context) error {
	req := req.ReqBookmark{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	err = r.GithubRepo.DelBookmark(
		c.Request().Context(),
		req.RepoName, claims.UserId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xoá bookmark thành công",
		Data:       nil,
	})
}