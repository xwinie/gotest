package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//分层的URL，Scheme后面紧跟着的是//
	u1, err := url.Parse("http://bing.com/search?q=dotnet#fragment")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("u1.Scheme", u1.Scheme) //http
	fmt.Println("u1.Opaque", u1.Opaque)
	fmt.Println("u1.User", u1.User)
	fmt.Println("u1.Host", u1.Host)         //bing.com
	fmt.Println("u1.Path", u1.Path)         ///search
	fmt.Println("u1.RawQuery", u1.RawQuery) //q=dotnet
	fmt.Println("u1.Fragment", u1.Fragment) //fragment

	fmt.Println()

	//透明类型的URL，Scheme后面不是//
	u2, err := url.Parse("mailto:xxx@163.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("u2.Scheme", u2.Scheme) //mailto
	fmt.Println("u2.Opaque", u2.Opaque) //xxx@163.com   适用于透明类型的URL
	fmt.Println("u2.User", u2.User)
	fmt.Println("u2.Host", u2.Host)
	fmt.Println("u2.Path", u2.Path)
	fmt.Println("u2.RawQuery", u2.RawQuery)
	fmt.Println("u2.Fragment", u2.Fragment)
}