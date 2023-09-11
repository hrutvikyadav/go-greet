package go_greet // notice the difference this is a so called shared library
// executable go programs have a main package instead of a named package

import "fmt"

func Greet(name string) string {
    // Function with Upcase 1st letter is callable from function in another package
    // Also called as exported name
    message := fmt.Sprintf("From %v motherfucker", name)
    return message
    
}

