check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./docs/swagger.json --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger -p 3540 swagger.json