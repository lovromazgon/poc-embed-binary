.PHONY: clean

build:
	go build -o entrypoint/main entrypoint/main.go
	go build -o plugin/main plugin/main.go
	go build main.go

clean:
	rm main
	rm -rf plugins
