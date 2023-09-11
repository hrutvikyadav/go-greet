package go_greet

import (
    "testing"
    "regexp"
)

// convention - TestSomeFunc() , takes in a pointer that can be used to report info from tests
// this one calls Greet with a name and check for correct results
func TestGreetName(t *testing.T) {
    // arrange
    name := "Hrx"
    wanted := regexp.MustCompile(`\b`+name+`\b`)
    // act
    message, err := Greet(name)
    // assert
    if !wanted.MatchString(message) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, wanted)
    }

}

// call Greet again to check for error
func TestHelloEmpty(t *testing.T) {
    msg, err := Greet("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
