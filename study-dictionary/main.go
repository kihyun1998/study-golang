package main

import (
	"fmt"
	"log"
	"study-golang/study-dictionary/dict"
)

func main() {
	dictionary := dict.Dictionary{}

	err := dictionary.Add("a", "is a")
	if err != nil {
		log.Fatalln(err)
	}

	err2 := dictionary.Update("a", "is b")
	if err != nil {
		log.Fatalln(err2)
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary.Search("a"))

	dictionary.Delete("a")
	fmt.Println(dictionary)
	fmt.Println(dictionary.Search("a"))
}
