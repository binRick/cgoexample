
all: clean build run

build:
	gcc -o libperson.so -Wall -g -shared -fPIC -lm person.c
	color black green
	echo BUILT libperson.so Shared Object
	color reset
	gcc -o hello -L. hello.c -lperson
	color black cyan
	echo BUILT hello Binary from C
	color reset
	env CGO_ENABLED=1 go build -o main main.go
	color black magenta
	echo BUILT main Binary from GO
	color reset

run:
	color black blue
	LD_LIBRARY_PATH=. ./hello
	color reset
	color black yellow
	./main
	color reset

clean:
	color black red 
	rm -rf hello libperson.so main
	echo DELETED hello, libperson.so, main
	color reset

