package main

import (
	"os"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/tax-calculator/database"
	"github.com/tax-calculator/middlewares"
	"github.com/urfave/cli"
	"github.com/gorilla/handlers"
)

func onError(err error, failedMessage string) {
	if err != nil {
		logrus.Errorln(failedMessage)
		logrus.Errorln(err)
	}
}

func runServer(db *gorm.DB) {
	r := LoadRouter(db)
	corsOption := middlewares.CorsMiddleware()
	logrus.Infoln("Server run on :8000")
	http.ListenAndServe(":8000", handlers.CORS(corsOption[0], corsOption[1], corsOption[2])(r))
}

func main() {
	db, err := database.InitDatabase()
	err = db.DB().Ping()

	if err != nil {
		logrus.Errorln(err)
	}

	defer db.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "Tax Calculator"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migration",
			Action: func(c *cli.Context) error {
				err = database.Migrate(db)
				onError(err, "Failed to migrate database schema")

				return err
			},
		},
		{
			Name:			"start",
			Description:	"Start REST API Server",
			Action:			func(c *cli.Context) error {
				runServer(db)
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		logrus.Fatalln(err)
	}
}