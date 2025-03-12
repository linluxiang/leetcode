import "strings"
import "fmt"

type Stack []byte

func (s *Stack) Len() int {
    return len(*s)
}

func (s *Stack) String() string {
    value := *s
    size := len(value)
    if size == 0 {
        return "0"
    }

    if size == 1 {
        return string(value[0])
    }

    v := strings.TrimLeft(string(value), "0")
    if v == "" {
        return "0"
    }
    return v
}

func (s *Stack) Pop() (byte, bool) {
    value := *s
    size := len(value)
    if size == 0 {
        return 0, false
    }
    last := value[size - 1]
    *s = value[:size - 1]
    return last, true
}

// Poping all the item less or equal than v until empty or poped k digits
func (s *Stack) PopUntil(v byte, k int) int {
    old := *s
    poped := 0
    for range k {
        last := old[len(old) - 1]
        if last <= v {
            break
        }
        // last >= v 
        old = old[:len(old) - 1]
        poped += 1
        if len(old) == 0 {
            break
        }
        // fmt.Println("pop: ", string(v), string(last), k)
    }
    *s = old
    return poped
}

func (s *Stack) Push(b byte) {
    *s = append(*s, b)
}

func (s *Stack) Top() (byte, bool) {
    value := *s
    size := len(value)
    if size == 0 {
        return 0, false
    }
    last := value[size - 1]
    return last, true
}

func removeKdigits(num string, k int) string {
    if k >= len(num) {
        return "0"
    }

    s := &Stack{}
    digits := []byte(num)

    removed := 0

    for _, d := range digits {
        if s.Len() == 0 || removed >= k {
            s.Push(d)
            continue
        }
        poped := s.PopUntil(d, k - removed)
        removed += poped
        s.Push(d)
        // fmt.Println("s: ", s.String())
    }

    if removed < k {
        for range k - removed{
            s.Pop()
        }
    }

    return s.String()
}
