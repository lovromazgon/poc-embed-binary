.PHONY: clean

build:
	go build -o main/main main/main.go
	go build -o embedded/main embedded/main.go
	go build combiner.go

clean:
	rm combiner
	rm -rf plugins
