build:
	go build -v

check: build ${GOPATH}/bin/errcheck
	go test -v
	${GOPATH}/bin/errcheck *.go

format:
	gofmt -w *.go

deps: ${GOPATH}/bin/errcheck ${GOPATH}/bin/dep
	${GOPATH}/bin/dep ensure

${GOPATH}/bin/errcheck:
	go get -u github.com/kisielk/errcheck

${GOPATH}/bin/dep:
	go get -u github.com/golang/dep/cmd/dep
