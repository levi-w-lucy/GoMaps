package main

import "fmt"

func main() {
	var google, aws string
	websites := map[*string]string{
		&google: "https://google.com",
		&aws:    "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites[&google])

	var linkedIn string
	websites[&linkedIn] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, &google)
	fmt.Println(websites)
}
