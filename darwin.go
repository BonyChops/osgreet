//go:build darwin

package osgreet

import "fmt"

func (*OSGreet) IsImplemented() bool {
	return true
}

func (*OSGreet) Greet() error {
	fmt.Println("Hello from macOS!")
	return nil
}
