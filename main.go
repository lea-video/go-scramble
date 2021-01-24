package main

import (
	"github.com/lea-video/go-scramble/scrambler"
)

func main() {
	println("generating seed: ")
	seed := scrambler.HashCode("test")
	println(seed)

	println("initialising")
	s, err := scrambler.New(scrambler.ConstructorOpts{Maximum: 15})
	panicOn(err)

	print("All next values: ")
	v := s.Next()
	for v != 0 {
		print(v, ", ")
		v = s.Next()
	}
	println()

	print("All prev values: ")
	v = s.Prev()
	for v != 0 {
		print(v, ", ")
		v = s.Prev()
	}
	println()
}

func panicOn(err error)  {
	if err != nil {
		panic(err)
	}
}
