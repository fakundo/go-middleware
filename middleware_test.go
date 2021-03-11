package middleware

import (
	"testing"
)

type originT = func(string) string

var origin = func(p string) string {
	return "origin:" + p
}

var md1 = Create(func(p string, next func()) string {
	next()
	return "md1:" + p
})

var md2 = Create(func(next func() string) string {
	return "md2:" + next()
})

func TestCreateMiddleware(t *testing.T) {
	dec1 := md1(origin)
	res1 := dec1.(originT)("param1")
	want1 := "md1:param1"
	if res1 != want1 {
		t.Errorf("got %q, want %q", res1, want1)
	}

	dec2 := md2(origin)
	res2 := dec2.(originT)("param2")
	want2 := "md2:origin:param2"
	if res2 != want2 {
		t.Errorf("got %q, want %q", res2, want2)
	}

	dec3 := md2(md1(origin))
	res3 := dec3.(originT)("param3")
	want3 := "md2:md1:param3"
	if res3 != want3 {
		t.Errorf("got %q, want %q", res3, want3)
	}
}

func TestUseMiddleware(t *testing.T) {
	dec := Use(md2, md1, origin)
	res := dec.(originT)("param3")
	want := "md2:md1:param3"
	if res != want {
		t.Errorf("got %q, want %q", res, want)
	}
}
