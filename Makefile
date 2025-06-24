.PHONY: run build start lin

run:
	go run main.go

build:
	go build -o build/main main.go

start: build
	./build/main


