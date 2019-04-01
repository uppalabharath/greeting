package greet

import (
	"fmt"
)

/*Greet greets the person. */
func Greet(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
