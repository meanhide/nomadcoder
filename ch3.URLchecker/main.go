package main

import "fmt"

func main() {
	urls := []string{
		"https://www.youtube.com/watch?v=iweGVkM9b_g",
		"https://www.youtube.com/watch?v=uvdC5FooVEU",
	}
	for _, url := range urls {
		fmt.Println(url)
	}
}
