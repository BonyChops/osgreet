//go:build darwin

package osgreet

import "fmt"

func Greet() {
	fmt.Println("Hello from macOS!")
}
