all: clean build

build:
	@go build -o ./build/cagt main.go &&\
	GOOS=windows GOARC=amd64 go build -o ./build/cagt.exe main.go

clean:
	@rm -rf ./build