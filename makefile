all: build

build:
	go build -o dagger main.go

run:
	./dagger

clean:
	rm -f ./dagger
