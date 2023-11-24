package main

import "math/rand"

func main() {

	type Month int

	const (
		January Month = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	for i := 0; i < 1000; i++ {
		print(months[rand.Intn(len(months))], "\n")
	}

}
