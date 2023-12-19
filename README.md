# Trie implementation in Go lang
A trie is a tree data structure used for storage and retrieval (hence the name 🌚) of strings. 
All strings stored in trie that have the same prefix also have common nodes.
The main feature of trie is the pattern matching in O(n) time, where n - length of the input.

## Example
```Go
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
```
## Output
```bash
$ make run
Has word address: true
Has prefix add:   true
Has word odd:     false
Now has word odd: true

Root:
-a
--d
---d
----r
-----e
------s
-------s
-o
--d
---d
```
