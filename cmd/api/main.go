package main

import (
	"flag"
	"github.com/BenefexLtd/departments-api-refactor/app"
	"github.com/BenefexLtd/onehub-go-base/pkg/config"
)

// main entry point of app

// load any config and start app...app configuration is defined in the app's app.go file.

func main() {

	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	config, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(app.Start(config))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
