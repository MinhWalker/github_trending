package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/helper"
	log "backend-github-trending/log"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

//call before main
func init() {
	fmt.Println("PRODUCTION ENVIROMENT")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	//connect db
	sql := &db.Sql{
		Host:     "35.221.222.42",	//localhost
		Port:     5432,
		UserName: "postgres",
		Password: "1234",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close() //exe func after defer when end main func

	//router
	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}

	api.SetupRouter()

	go scheduleUpdateTrending(15 * time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for{
			select {
			case <- ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}