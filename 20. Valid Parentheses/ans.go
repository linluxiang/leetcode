type Stack struct {
    Content []rune
}

func (s *Stack) Top() rune {
    if len(s.Content) == 0 {
        return 0
    }
    return s.Content[len(s.Content) - 1]
} 

func (s *Stack) Pop() rune {
    if len(s.Content) == 0 {
        return 0
    }
    res := s.Content[len(s.Content) - 1]
    s.Content = s.Content[0:len(s.Content) - 1]
    return res
}

func (s *Stack) Push(v rune) {
    s.Content = append(s.Content, v)
}

func isValid(s string) bool {
	if len(s) % 2 == 1 {
		return false
    }

    stack := &Stack{}
    for _, c := range s {
        top := stack.Top()
        if top == 0 {
            stack.Push(c)
            continue
        }
        if c == '(' || c== '{' || c == '[' {
            stack.Push(c)
            continue
        }

        if c == ')' && top != '(' {
            return false
        }

        if c == ']' && top != '[' {
            return false
        }

        if c == '}' && top != '{' {
            return false
        }
        // pair successfully

        stack.Pop()

    }
    return len(stack.Content) == 0
}
