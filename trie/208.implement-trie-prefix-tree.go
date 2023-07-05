package main

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	nextChars [26]*Trie
	isWord    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curNode := this
	for i := 0; i < len(word); i++ {
		nextNode := curNode.nextChars[word[i]-'a']
		if nextNode == nil {
			nextNode = &Trie{}
			curNode.nextChars[word[i]-'a'] = nextNode
		}
		curNode = nextNode
	}
	curNode.isWord = true
}

func (this *Trie) Search(word string) bool {
	curNode := this
	for i := 0; i < len(word); i++ {
		nextNode := curNode.nextChars[word[i]-'a']
		if nextNode == nil {
			return false
		}
		curNode = nextNode
	}
	return curNode.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curNode := this
	for i := 0; i < len(prefix); i++ {
		nextNode := curNode.nextChars[prefix[i]-'a']
		if nextNode == nil {
			return false
		}
		curNode = nextNode
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
// @lc code=end
