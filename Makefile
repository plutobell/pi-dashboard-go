# @Program : Pi Dashboard Go (https://github.com/plutobell/pi-dashboard-go)
# @Description: Golang implementation of pi-dashboard
# @Author: github.com/plutobell
# @Creation: 2020-8-10
# @Last modify: 2020-8-10
# @Version: 1.0.8

PROGRAM = pi-dashboard-go
OUTPUT = build
GOOS = linux

build: clean rice-box.go main.go server.go device.go  device_test.go go.mod go.sum
	@echo "Building"

	@echo "1 -> Building the "${PROGRAM}_${GOOS}_arm
	@GOOS=${GOOS} GOARCH=arm GOARM=6 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_arm

	@echo "2 -> Building the "${PROGRAM}_${GOOS}_armv5
	@GOOS=${GOOS} GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv5

	@echo "3 -> Building the "${PROGRAM}_${GOOS}_armv6
	@GOOS=${GOOS} GOARCH=arm GOARM=6 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv6

	@echo "4 -> Building the "${PROGRAM}_${GOOS}_armv7
	@GOOS=${GOOS} GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_armv7

	@echo "5 -> Building the "${PROGRAM}_${GOOS}_arm64
	@GOOS=${GOOS} GOARCH=arm64 GOARM=7 go build -ldflags "-s -w" -o ./${OUTPUT}/${PROGRAM}_${GOOS}_arm64

	@echo "Complete"

run:
	@echo "Running"
	@go run ./
	@echo "Complete"

test:
	@echo "Testing"
	@go test -v
	@go test -test.bench=".*"
	@echo "Complete"

clean:
	@echo "Cleaning"
	@rm -rf ./build
	@echo "Complete"

rice-box.go:
	@echo "Generate rice-box.go"
	@apt install golang-rice -y > /dev/null 2> /dev/null
	@rice embed-go
	@echo "Complete"

help:
	@echo "Commands: build | run | test | clean"