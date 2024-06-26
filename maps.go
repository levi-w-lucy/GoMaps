package main

import "fmt"

type floatMap = map[string]float64

func main() {
	userNames := make([]string, 0, 5)

	userNames = append(userNames, "max")
	userNames = append(userNames, "manuel")
	fmt.Println(userNames)

	courseRatings := make(floatMap)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	loop(courseRatings)
	fmt.Println(courseRatings)

}

func loop[T comparable, K any](itemToLoop map[T]K) {
	for index, value := range itemToLoop {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}
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
