package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// a url ncludes a scheme, authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// scheme
	fmt.Println(u.Scheme)
	// auth info
	// User is sort of key-value pair, *url.Userinfo
	fmt.Println(u.User)
	fmt.Printf("u.User is of %T\n", u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// host and port
	fmt.Println(u.Host)
	// u.Host is string
	fmt.Printf("u.Host is of %T\n", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// path, also string
	fmt.Printf("u.Path is of %T\n", u.Path)
	fmt.Println(u.Path)
	// what is a fragment?
	fmt.Println(u.Fragment)

	// query, also string
	fmt.Println(u.RawQuery)
	fmt.Printf("u.RawQuery is of %T\n", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}