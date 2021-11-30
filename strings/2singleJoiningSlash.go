package main

import "strings"

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func main() {
	println(singleJoiningSlash("foo", "bar"))
	println(singleJoiningSlash("foo/", "bar"))
	println(singleJoiningSlash("foo", "/bar"))
	println(singleJoiningSlash("foo/", "/bar"))
}
