package main

import (
	"fmt"

	"github.com/Adrone13/trie/src/trie"
)

func main() {
	t := trie.New()

	t.Insert("address")
	fmt.Println("Has word address:", t.Search("address"))
	fmt.Println("Has prefix add:  ", t.StartsWith("add"))
	fmt.Println("Has word odd:    ", t.Search("odd"))
	t.Insert("odd")
	fmt.Println("Now has word odd:", t.Search("odd"))

	fmt.Println()

	t.Print()
}
