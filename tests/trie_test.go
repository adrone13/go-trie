package tests

import (
	"testing"

	"github.com/Adrone13/trie/src/trie"
)

func TestTrie(t *testing.T) {
	trie := trie.New()

	trie.Insert("abc")
	if trie.Search("abc") != true {
		t.Errorf("expected to return true")
	}
	if trie.Search("xyz") != false {
		t.Errorf("expected to return false")
	}
	if trie.Search("ab") != false {
		t.Errorf("expected to return true")
	}
	if trie.StartsWith("ab") != true {
		t.Errorf("expected to return true")
	}
	if trie.StartsWith("abc") != true {
		t.Errorf("expected to return true")
	}
	if trie.StartsWith("xyz") != false {
		t.Errorf("expected to return true")
	}

	// if resp.StatusCode != http.StatusOK {
	// 	t.Errorf("expected status OK; got %v", resp.Status)
	// }
}
