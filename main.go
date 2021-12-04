package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lperson
#include "person.h"
*/
import "C"
import (
	"fmt"
	"os"
)

type (
	Person C.struct_APerson
)

func (p *Person) Name() string {
	return C.GoString(p.name)
}

func (p *Person) LongName() string {
	return C.GoString(p.long_name)
}

func GetPerson(name string, long_name string) *Person {
	return (*Person)(C.get_person(C.CString(name), C.CString(long_name)))
}

func main() {
	fmt.Fprintf(os.Stderr, "GO MAIN START\n")
	var f *Person
	f = GetPerson("tim", "tim hughes")
	fmt.Printf("C BINDING>          Name: %s | Long Name %s | \n", C.GoString(f.name), C.GoString(f.long_name))

	fmt.Printf("GO BINDING>         Name: %s | Long Name:  %s | \n", f.Name(), f.LongName())
	fmt.Fprintf(os.Stderr, "GO MAIN END\n")
}
