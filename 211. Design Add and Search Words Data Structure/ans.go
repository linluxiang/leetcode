type Node struct {
    children [26]*Node
    isEnd bool
}

func (n *Node) search(word []byte) bool {
    if len(word) == 0 {
        return n.isEnd
    }

    if word[0] == '.' {
        for i := range 26 {
            if n.children[i] != nil {
                result := n.children[i].search(word[1:len(word)])
                if result == true {
                    return true
                }
            }
        }
    } else {
        i := word[0] - 'a'
        if n.children[i] != nil {
            return n.children[i].search(word[1:len(word)])
        }
    }

    return false
}

type WordDictionary struct {
    root *Node
}

func NewNode() *Node {
    return &Node{
        children: [26]*Node{},
        isEnd: false,
    }
}

func Constructor() WordDictionary {
    return WordDictionary {
        root: NewNode(),
    }
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.root
    for _, c := range []byte(word) {
        idx := c - 'a'

        if node.children[idx] == nil {
            node.children[idx] = NewNode()
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

