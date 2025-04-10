package redacted_test

import (
	"fmt"

	"github.com/andrewheberle/redacted-string"
)

func ExampleRedact() {
	s := "this is a string"
	fmt.Println(redacted.Redact(s))
	// Output: ********
}

func ExampleNew() {
	s := "this is a string"

	censor := redacted.New(s)
	fmt.Println(censor)
	// Output: ********
}

func ExampleCensor_Set() {
	s := "this is a string"

	censor := redacted.New(s, redacted.WithLength(0))
	censor.Set("a new string that we want to be fully replaced with the * character")
	fmt.Println(censor)
	// Output: *******************************************************************
}

func ExampleWithLength() {
	s := "this is a string"

	censor := redacted.New(s, redacted.WithLength(12))
	fmt.Println(censor)
	// Output: ************
}

func ExampleWithCharacter() {
	s := "this is a string"

	censor := redacted.New(s, redacted.WithCharacter('.'))
	fmt.Println(censor)
	// Output: ........
}
