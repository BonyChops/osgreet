//go:build windows

package osgreet

import "fmt"

func (*OSGreet) IsImplemented() bool {
	return true
}

func (*OSGreet) Greet() error {
	fmt.Println("Hello from Windows!")
	return nil
}
