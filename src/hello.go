package main

import "fmt"

const hello_and_a_comma = "Hello, "

func Hello(name string) string {
    return hello_and_a_comma+name
}

func main() {
    fmt.Println(Hello("world"))
}