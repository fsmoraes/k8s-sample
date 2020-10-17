package main

import "fmt"

func greeting(message string) string {
	return fmt.Sprintf("<b>%v</b>", message)
}

func main() {
	fmt.Println(greeting("hello"))
}
