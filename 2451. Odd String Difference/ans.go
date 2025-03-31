func difference(word string) []byte {
    result := make([]byte, len(word) - 1)
    for i := range len(word) - 1 {
        result[i] = word[i+1] - word[i]
    }

    return result
}

func equal(a, b []byte) bool {
    for i := range len(a) {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func oddString(words []string) string {
    diff0 := difference(words[0])
    diff1 := difference(words[1])

    // diff0Cnt := 1
    // diff1Cnt := 1
    equal01 := equal(diff0, diff1)

    // ans := words[0]
    for i := 2; i < len(words); i++ {
        if equal01 {
            if !equal(diff0, difference(words[i])) {
                return words[i]
            }
        } else {
        if equal(diff0, difference(words[i])) {
            return words[1]
        } else {
            return words[0]
        }
        }
    }
    // return ans
    return ""
}
