import "fmt"

type Trie struct {
	Val      string
	IsEnd    bool
	Children []*Trie
}

func Constructor() Trie {
	return Trie{
		Val:      "",
		IsEnd:    true,
		Children: []*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
        this.IsEnd = true
		return
	}

	for _, child := range this.Children {
		if child.Val == string(word[0]) {
			child.Insert(word[1:len(word)])
			return
		}
	}

	this.Children = append(this.Children, &Trie{
		Val:      string(word[0]),
		IsEnd:    false,
		Children: []*Trie{},
	})

    this.Children[len(this.Children) - 1].Insert(word[1:len(word)])

}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.IsEnd
	}

	for _, child := range this.Children {
		if child.Val == string(word[0]) {
			return child.Search(word[1:len(word)])
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	for _, child := range this.Children {
		if child.Val == string(prefix[0]) {
			return child.StartsWith(prefix[1:len(prefix)])
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
