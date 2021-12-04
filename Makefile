
all: clean build run

build:
	gcc -o libperson.so -Wall -g -shared -fPIC -lm person.c
	gcc -o hello -L. hello.c -lperson
	env CGO_ENABLED=1 go build -o main main.go

run:
	LD_LIBRARY_PATH=. ./hello
	./main

clean:
	rm -rf hello libperson.so main

