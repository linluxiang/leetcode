import "strings"

type Stack struct {
    stack []string
}

func (this *Stack) Push(v string) {
    this.stack = append(this.stack, v)
}

func (this *Stack) Pop() string {
    if len(this.stack) == 0 {
        return ""
    }

    last := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    return last
}

func (this *Stack) Top() string {
    if len(this.stack) == 0 {
        return ""
    }

    return this.stack[len(this.stack) - 1]
}

func (this *Stack) Output() string {
    root := "/"
    root += strings.Join(this.stack, "/")
    return root
}

func findNextWord(path string, i int) (string, int) {
    pos := i + 1
    word := ""
    for pos < len(path) && path[pos] != '/' {
        word += string(path[pos])
        pos += 1
    }
    return word, pos
}

func simplifyPath(path string) string {
    i := 0
    stack := &Stack{}
    var word string
    for i < len(path) {
        word, i = findNextWord(path, i)
        if word == "" {
            // Do nothing
        } else if word == "." {
            // Do nothing
        } else if word == ".." {
            stack.Pop()
        } else {
            stack.Push(word)
        }

        if i == len(path) {
            break
        }
    }

    return stack.Output()
}
