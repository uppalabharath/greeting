package greet

import (
	"errors"
	"fmt"
)

/*Greet greets the person. */
func Greet(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "es":
		return fmt.Sprintf("Hola, %s!", name), nil
	default:
		return "", errors.New("Unknown Language")
	}
}
