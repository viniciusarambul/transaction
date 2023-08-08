package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/viniciusarambul/transaction/db"
	"github.com/viniciusarambul/transaction/handler"
	"github.com/viniciusarambul/transaction/log"
	"github.com/viniciusarambul/transaction/presenter"
	"github.com/viniciusarambul/transaction/repository"
	"github.com/viniciusarambul/transaction/usecase"
	"github.com/viniciusarambul/transaction/utils"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

const (
	indexTmpl   = "index"
	profileTmpl = "profile"
)

func main() {
	ctx := context.Background()
	engine := gin.Default()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	engine.Use(otelgin.Middleware("ms-transaction"))

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})

	})

	log, err := log.InitLogger()
	if err != nil {
		fmt.Println(err)
	}

	db, err := db.SetupDB()
	if err != nil {
		fmt.Println(err)
		panic("errou")
	}

	log.Info("Init database successful")

	accountRepository := repository.NewAccountRepository(db)
	accountPresenter := presenter.NewAccountPresenter()
	accountUseCase := usecase.NewAccountUseCase(accountRepository, accountPresenter, log)

	handler.NewAccountHandler(engine, accountUseCase)

	clock := utils.New()
	transactionRepository := repository.NewTransactionRepository(db)
	operationRepository := repository.NewOperationRepository(db)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository, log, clock)

	handler.NewTransactionHandler(engine, transactionUseCase)

	engine.Run()
}

func parseTemplates() *template.Template {
	indexTemplate := `
		<html>
		<p>Here are some routes for you:</p>
		<ul>
			<li><a href="/hello/world">Hello world</a></li>
			<li><a href="/hello/foo-bar">Hello foo-bar</a></li>
		</ul>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	t := template.Must(template.New(indexTmpl).Parse(indexTemplate))

	profileTemplate := `
		<html>
		<h3>Hello {{ .username }}</h3>
		<p><a href="{{ .traceURL }}" target="_blank">{{ .traceURL }}</a></p>
		</html>
	`
	return template.Must(t.New(profileTmpl).Parse(profileTemplate))
}
