type Set struct {
    set map[rune]int
    Values []string
}

func NewSet(v string) *Set{
    res := &Set {
        set: make(map[rune]int, len(v)),
        Values: []string {v},
    }
    for _, c := range v {
        res.set[c] += 1
    }
    return res
}

func (s *Set) Equal(v string) bool {

    count := map[rune]int{}

    for _, c := range v {
        if _, ok := s.set[c]; !ok {
            return false
        }
        count[c] += 1
    }

    for _, c := range v {
        if count[c] != s.set[c] {
            return false
        }
    }

    if len(count) != len(s.set) {
        return false
    }
    return true
}

func (s *Set) Add(v string) {
    s.Values = append(s.Values, v)
}


func groupAnagrams(strs []string) [][]string {
    var sets = make([]*Set, 0, len(strs))

outer:
    for _, str := range strs {
        for _, set := range sets {
            if set.Equal(str) {
                set.Add(str)
                continue outer
            }
        }
        // No set contains this string
        newSet := NewSet(str)
        sets = append(sets, newSet)
    }

    var res = make([][]string, 0, len(strs))

    for _, set := range sets {
        res = append(res, set.Values)
    }

    return res
}
