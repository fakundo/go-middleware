package middleware

import (
	"fmt"
	"testing"
)

var origin = func(a string) string {
	fmt.Println("origin", a)
	return "origin result" + a
}

var md1 = Create(func(next func() string) string {
	fmt.Println("md1")
	return next()
	// return "md1 result"
})

func TestCreateMiddleware(t *testing.T) {
	decorated := md1(origin)
	res := decorated.(func(string) string)("qwe")
	fmt.Println(res)

	// want := "Hello, world."
	// got :=

	// want := "Hello, world."
	// if got := CreateMiddleware(func() {}); got != want {
	// 	t.Errorf("Hello() = %q, want %q", got, want)
	// }
}

func TestUseMiddleware(t *testing.T) {

}
