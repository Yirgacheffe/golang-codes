package main

import "fmt"

// Foo server has many parameters
type Foo struct {
	verbosity int
	logging   bool
}

type option func(*Foo)

func (f *Foo) Option(opts ...option) {
	for _, opt := range opts {
		opt(f)
	}
}

// option 1
func Verbosity(v int) option {
	return func(f *Foo) {
		f.verbosity = v
	}
}

func Loggin(enable bool) option {
	return func(f *Foo) {
		f.logging = enable
	}
}

// type option func(*Foo) interface{}
/*
func Verbosity(v int) option {
	return func(f *Foo) interface{} {
		prev := f.verbosity
		f.verbosity = v
		return prev
	}
}
*/

func main() {
	f := &Foo{}
	f.Option(Verbosity(3), Loggin(false))

	fmt.Println(f.verbosity, f.logging)

	/*
		prevVerbosity := foo.Option(pkg.Verbosity(3))
		foo.DoSomeDebugging()
		foo.Option(pkg.Verbosity(prevVerbosity.(int)))
	*/
}
