# rest-api
Sample Golang Rest API structure

based on the gorsk template (https://github.com/ribice/gorsk)

Project Structure
Root directory contains things not related to code directly, e.g. docker-compose, CI/CD, readme, bash scripts etc. It should also contain vendor folder, Gopkg.toml and Gopkg.lock if dep is being used.

Cmd package contains code for starting applications (main packages). The directory name for each application should match the name of the executable you want to have. Gorsk is structured as a monolith application but can be easily restructured to contain multiple microservices. An application may produce multiple binaries, therefore Gorsk uses the Go convention of placing main package as a subdirectory of the cmd package. As an example, scheduler application's binary would be located under cmd/cron. It also loads the necessery configuration and passes it to the service initializers.

Rest of the code is located under /pkg. The pkg directory contains utl and 'microservice' directories.

Microservice directories, like api (naming corresponds to cmd/ folder naming) contains multiple folders for each domain it interacts with, for example: user, car, appointment etc.

Domain directories, like user, contain all application/business logic and two additional directories: platform and transport.

Platform folder contains various packages that provide support for things like databases, authentication or even marshaling. Most of the packages located under platform are decoupled by using interfaces. Every platform has its own package, for example, postgres, elastic, redis, memcache etc.

Transport package contains HTTP handlers. The package receives the requests, marshals, validates then passes it to the corresponding service.

Utl directory contains helper packages and models. Packages such as mock, middleware, configuration, server are located here.
