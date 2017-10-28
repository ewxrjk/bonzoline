build:
	go build

check: build
	go test
	${GOPATH}/bin/errcheck *.go

format:
	gofmt -w *.go

deps: ${GOPATH}/bin/errcheck 
	${GOPATH}/bin/dep ensure

${GOPATH}/bin/errcheck:
	go get -u github.com/kisielk/errcheck
