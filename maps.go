package main

import "fmt"

func main() {
	userNames := make([]string, 0, 5)

	fmt.Println(len(userNames), cap(userNames))
	userNames = append(userNames, "max")
	userNames = append(userNames, "manuel")
	fmt.Println(userNames)

	fmt.Println(len(userNames), cap(userNames))
	fmt.Println(userNames)

}

func mapTesting() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
