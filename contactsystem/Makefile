override BINARY=contactsystem
VERSION=V1.0.0
API_ENV?=local
host?="localhost"
port?="8080"

compile: ## Compile the project to generate the binary in the target folder
ifndef VERSION
	$(error VERSION environment variable has to be set)
endif
	(cd cmd/http/ && go build -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=`git rev-parse HEAD` -X main.BuildTime=`date +%FT%T%z`" -o ../../build/${BINARY})


run: ## executes the invoice_system service
	cd build;./${BINARY} 