# @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
# @Description: Golang implementation of pi-dashboard
# @Author: github.com/plutobell
# @Creation: 2020-08-10
# @Last modification: 2021-08-13
# @Version: 1.3.2

PROGRAM = pi-dashboard-go
OUTPUT = build
GOOS = linux
OS_NAME = $(shell uname -o)

build: clean vet main.go server.go device.go  device_test.go go.mod go.sum
	@echo "-> Building"

	@echo "-> 1 Building the "${PROGRAM}_${GOOS}_armv5_32
	@GOOS=${GOOS} GOARCH=arm GOARM=5 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv5_32

	@echo "-> 2 Building the "${PROGRAM}_${GOOS}_armv6_32
	@GOOS=${GOOS} GOARCH=arm GOARM=6 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv6_32

	@echo "-> 3 Building the "${PROGRAM}_${GOOS}_armv7_32
	@GOOS=${GOOS} GOARCH=arm GOARM=7 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv7_32

	@echo "-> 4 Building the "${PROGRAM}_${GOOS}_armv5_64
	@GOOS=${GOOS} GOARCH=arm64 GOARM=5 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv5_64

	@echo "-> 5 Building the "${PROGRAM}_${GOOS}_armv6_64
	@GOOS=${GOOS} GOARCH=arm64 GOARM=6 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv6_64

	@echo "-> 6 Building the "${PROGRAM}_${GOOS}_armv7_64
	@GOOS=${GOOS} GOARCH=arm64 GOARM=7 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv7_64

	@echo "-> 7 Building the "${PROGRAM}_${GOOS}_386
	@GOOS=${GOOS} GOARCH=386 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_386

	@echo "-> 8 Building the "${PROGRAM}_${GOOS}_amd64
	@GOOS=${GOOS} GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_amd64

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
	@rm -rf ./build
	@echo "-> Complete"

help:
	@echo "-> Commands: build | run | vet | test | clean | help"