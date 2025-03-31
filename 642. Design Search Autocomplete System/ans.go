import "slices"
import "strings"
import "fmt"

type Node struct {
    children [27]*Node
    value string
}

func (n *Node) Insert(v string) {
    n.insert(v, v)
}

func (n *Node) insert(v string, raw string) {
    // fmt.Println("insert: ", v)

    if len(v) == 0 {
        n.value = raw
        return
    }
    c := v[0]
    idx := c - 'a'
    if c == ' ' {
        idx = 26
    }

    if n.children[idx] == nil {
        n.children[idx] = &Node{}
    }

    n.children[idx].insert(v[1:len(v)], raw)
}

func (n *Node) GetAllSubValues(pattern []byte) []string {
    ans := []string{}
    n.getAllSubValues(pattern, &ans)
    return ans
}

func (n *Node) getAllSubValues(pattern []byte, ans *[]string) {
    // fmt.Println("getAllSub: '", string(pattern), "',", *ans)
    if len(pattern) != 0 {
        c := pattern[0]
        idx := c - 'a'
        if c == ' ' {
            idx = 26
        }
        // fmt.Println("searching for c: ", string(c))
        if n.children[idx] == nil {
            return
        }
        n.children[idx].getAllSubValues(pattern[1:len(pattern)], ans)
        return
    }
    if len(n.value) != 0 {
        *ans = append(*ans, n.value)
    }
    for i := 0; i < 27; i++ {
        if n.children[i] != nil {
            n.children[i].getAllSubValues(pattern, ans)
        }
    }
    // fmt.Println("ans: ", *ans)
}

type AutocompleteSystem struct {
    trie *Node
    appears map[string]int
    pattern []byte
}


func Constructor(sentences []string, times []int) AutocompleteSystem {
    t := &Node{}
    m := map[string]int{}

    for i := range sentences {
        t.Insert(sentences[i])
        m[sentences[i]] = times[i]
    }

    return AutocompleteSystem {
        trie: t,
        appears: m,
        pattern: []byte{},
    }
}


func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        v := string(this.pattern)
        this.trie.Insert(v)
        this.appears[v] += 1
        this.pattern = this.pattern[:0]
        return []string{}
    }

    this.pattern = append(this.pattern, c)
    // fmt.Printf("this.pattern: %v\n", this.pattern)
    values := this.trie.GetAllSubValues(this.pattern)
    slices.SortFunc(values, func(a, b string) int {
        if this.appears[b] != this.appears[a] {
            return this.appears[b] - this.appears[a]
        }
        return strings.Compare(a, b)
    })

    // fmt.Println(this.pattern, ",values: ", values, this.appears)

    size := min(3, len(values))
    return values[:size]
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */
