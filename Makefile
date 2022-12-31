## start migration
export


help:  ## show this help
	@echo "usage: make [target]"
	@echo ""
	@egrep "^(.+)\:\ .*##\ (.+)" ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

run: ## run it will instance server 
	go run cmd/app.go

.PHONY: test/cov
test/cov:
	go test --cover -coverpkg=./...  ./... -coverprofile=cover_app.out
	go tool cover -html=cover_app.out