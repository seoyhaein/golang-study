package main

import "fmt"

// Fooer has to Foo
type Fooer interface {
	Foo()
}

// Bar is a proxy, that calls Foo of specific instance.
func Bar(a Fooer) {
	a.Foo()
}

//////////////////////////////////////////////////////////////////////
// usage

func main() {
	a := &A{} // note it is a pointer
	// also there's no need to specify values for default-initialized fields.
	Bar(a)
}

//////////////////////////////////////////////////////////////////////
// implementation

// A is a "base class"
type A struct {
	unimplementB
}

//func (a *A) Foo() {
//    fmt.Println("Hello World")
//}

type unimplementB struct {
}

func (b *unimplementB) Foo() {
	fmt.Println("not implemented")
}
