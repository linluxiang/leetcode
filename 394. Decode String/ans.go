import "fmt"


type Stack []string

func (s *Stack) Push(str string) {
    *s = append(*s, str)
}

func (s *Stack) Pop() string {
    size := len(*s)
    if size == 0 {
        return ""
    }

    last := (*s)[size - 1]
    *s = (*s)[:size-1]
    return last
}


func isDigit(c byte) bool {
    return '0' <= c && c <= '9'
}

func isLetter(c byte) bool {
    return 'a' <= c && c <= 'z'
}

func decodeString(s string) string {
    stack := &Stack{}
    tempC := 0
    for _, c := range []byte(s) {
        // fmt.Println(*stack)
        if isDigit(c) {
            tempC = tempC*10 + int(c - '0')
            continue
        }
        if c == '[' {
            v := strconv.Itoa(tempC)
            stack.Push(v)
            tempC = 0
            continue
        }
        if isLetter(c) {
            stack.Push(string(c))
            continue
        }

        if c == ']' {
            tempS := ""
            inner:
            for {
                last := stack.Pop()
                if isDigit(last[0]) {
                    if tempS != "" {
                        cnt, _ := strconv.Atoi(last)
                        // fmt.Println("cnt: ", cnt)
                        v := ""
                        for range cnt {
                            v += tempS
                        }
                        // fmt.Println("cnt: ", cnt, ", v: ", v)
                        stack.Push(v)
                        
                    }
                    break inner
                }
                tempS = last + tempS
                // fmt.Println("temp: ", tempS)
            }
        }
    }
    result := ""
    v := stack.Pop()
    for v != "" {
        result = v + result
        v = stack.Pop()
    }
    return result
}
