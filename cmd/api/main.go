package main
// main entry point of app

// load any config and start app...app configuration is defined in the app's app.go file.


/*
Cmd package contains code for starting applications (main packages).
The directory name for each application should match the name of the executable you want to have.
Gorsk is structured as a monolith application but can be easily restructured to contain multiple microservices.
An application may produce multiple binaries, therefore Gorsk uses the Go convention of placing main package as a subdirectory of the cmd package.
As an example, scheduler application's binary would be located under cmd/cron. It also loads the necessery configuration and passes it to the service initializers.
*/

import (
	"flag"
	"github.com/BenefexLtd/departments-api-refactor/app"

	"github.com/BenefexLtd/departments-api-refactor/app/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./cmd/pkg/conf.local.yaml", "Path to config file")
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