package main

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRun(t *testing.T) {
    // name := "Gladys"
    // want := regexp.MustCompile(`\b`+name+`\b`)
    // msg, err := Hello("Gladys")
    // if !want.MatchString(msg) || err != nil {
    //     t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    // }

	// input = [ '', '', '' ]

	// app, err = cli.Run(input)
	// if (app.____  != _____ ||
		// 	app._____ != ""
		// ) {
	// 	f.Fatalf("App configuration incorrect", app, err)
	// }

}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestHelloEmpty(t *testing.T) {
//     msg, err := Hello("")
//     if msg != "" || err == nil {
//         t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//     }
// }