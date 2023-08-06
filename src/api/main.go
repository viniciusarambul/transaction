package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/viniciusarambul/transaction/src/api/handler"
	"github.com/viniciusarambul/transaction/src/api/presenter"
	"github.com/viniciusarambul/transaction/src/infra"
	"github.com/viniciusarambul/transaction/src/infra/repository"
	"github.com/viniciusarambul/transaction/src/usecase"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	db, err := infra.SetupDB()
	if err != nil {
		fmt.Println(err)
		panic("errou")
	}

	accountRepository := repository.NewAccountRepository(db)
	accountPresenter := presenter.NewAccountPresenter()
	accountUseCase := usecase.NewAccountUseCase(accountRepository, accountPresenter)

	handler.NewAccountHandler(engine, accountUseCase)

	engine.Run()
}
