# @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
# @Description: Golang implementation of pi-dashboard
# @Author: github.com/plutobell
# @Creation: 2020-8-10
# @Last modify: 2020-8-14
# @Version: 1.0.9

PROGRAM = pi-dashboard-go
OUTPUT = build
GOOS = linux
OS_NAME = $(shell uname -o)

build: clean vet rice-box.go main.go server.go device.go  device_test.go go.mod go.sum
	@echo "-> Building"

	@echo "-> 1 Building the "${PROGRAM}_${GOOS}_arm
	@GOOS=${GOOS} GOARCH=arm GOARM=6 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_arm

	@echo "-> 2 Building the "${PROGRAM}_${GOOS}_armv5
	@GOOS=${GOOS} GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv5

	@echo "-> 3 Building the "${PROGRAM}_${GOOS}_armv6
	@GOOS=${GOOS} GOARCH=arm GOARM=6 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv6

	@echo "-> 4 Building the "${PROGRAM}_${GOOS}_armv7
	@GOOS=${GOOS} GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv7

	@echo "-> 5 Building the "${PROGRAM}_${GOOS}_arm64
	@GOOS=${GOOS} GOARCH=arm64 GOARM=7 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_arm64

	@echo "-> 6 Building the "${PROGRAM}_${GOOS}_386
	@GOOS=${GOOS} GOARCH=386 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_386

	@echo "-> 7 Building the "${PROGRAM}_${GOOS}_amd64
	@GOOS=${GOOS} GOARCH=amd64 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_amd64

	@echo "-> Complete"

run: clean vet
	@echo "-> Running"
	@go run ./
	@echo "-> Complete"

vet:
	@echo "-> Checking"
	@go vet
	@echo "-> Complete"

test:
	@echo "-> Testing"
	@go test -v
	@go test -test.bench=".*"
	@echo "-> Complete"

clean:
	@echo "-> Cleaning"
	@rm -rf rice-box.go
	@rm -rf ./build
	@echo "-> Complete"

rice-box.go:
	@echo "-> Generate rice-box.go"
ifeq ($(OS_NAME), GNU/Linux)
	@apt install golang-rice -y > /dev/null 2> /dev/null
else
	@go get github.com/GeertJohan/go.rice > /dev/null 2> /dev/null
	@go get github.com/GeertJohan/go.rice/rice > /dev/null 2> /dev/null
endif
	@rice embed-go
	@echo "-> Complete"

help:
	@echo "-> Commands: build | run | test | vet | clean | help"