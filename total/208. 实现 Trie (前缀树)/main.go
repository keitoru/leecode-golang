package main

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range []byte(word) {
		if node.next[c-'a'] == nil {
			node.next[c-'a'] = &Trie{}
		}
		node = node.next[c-'a']
	}

	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range []byte(word) {
		node = node.next[c-'a']
		if node == nil {
			return false
		}
	}

	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range []byte(prefix) {
		node = node.next[c-'a']
		if node == nil {
			return false
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
