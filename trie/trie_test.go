package trie

import (
	"reflect"
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")
	testFind(t, trie, "fo", []string{"foo", "fool", "foolish"})

	trie.Delete("foolis")
	testFind(t, trie, "fo", []string{"foo", "fool", "foolish"})

	trie.Delete("fook")
	testFind(t, trie, "fo", []string{"foo", "fool", "foolish"})

	trie.Delete("fool")
	testFind(t, trie, "fo", []string{"foo", "foolish"})

	trie.Insert("fool")
	trie.Delete("foolish")
	testFind(t, trie, "fo", []string{"foo", "fool"})
}

func testFind(t *testing.T, trie *Trie, search string, expected []string) {
	got := trie.Find(search)
	sort.Strings(got)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Find(\"%s\") = %v; want %v", search, got, expected)
	}
}
