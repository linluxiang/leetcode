type Node struct {
    Val string
    Child *Node
}

func Insert(node *Node, value string) {
    if len(value) == 0 {
        return
    }
    val := string(value[0])

    node.Child = &Node{val, nil}
    if len(value) >1 {
        Insert(node.Child, value[1:len(value)])
    }
}

func Match(node *Node, value string) {
    if len(value) == 0 {
        if node.Child != nil {
            node.Child = nil
        }
        return
    }
    if node.Child == nil {
        return
    }

    val := string(value[0])

    if node.Child.Val != val {
        node.Child = nil
        return
    }


    Match(node.Child, value[1:len(value)])
}

func BuildString(node *Node) string {
    if node.Child == nil {
        return node.Val
    }
    return node.Val + BuildString(node.Child)
}

func longestCommonPrefix(strs []string) string {
    root := &Node{"", nil}
    for i, str := range strs {
        if i == 0 {
            Insert(root, str)
        } else {
            Match(root, str)
        }

    }

    return BuildString(root)
}
