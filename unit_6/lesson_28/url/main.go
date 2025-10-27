package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	url, err := url.Parse("https://a b.com")
	if err != nil {
		fmt.Printf("%#v\n", err)
		if err, ok := err.(net.Error); ok {
			fmt.Println(err.Error())
			fmt.Println(err.Timeout())
		}
		panic(err)
	}
	fmt.Println(url)
}
