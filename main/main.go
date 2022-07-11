package main

import "fmt"

const helloEnglishPrefix = "Hello, "

func Hello(name string) string {
	return helloEnglishPrefix + name
}

func main() {
	fmt.Println(Hello("Umar"))
}
