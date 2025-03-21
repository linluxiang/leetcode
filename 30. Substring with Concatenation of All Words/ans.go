func isStringConcatenated(s string, wordLen int, words map[string]int) bool {
    if len(s) == 0 {
        return false
    }
    appears := map[string]int{}
    
    for k := range words {
        appears[k] = 0
    }

    // fmt.Println("checking: ", s)
    for i := 0; i < len(s); i += wordLen {
        sub := s[i:i+wordLen]
        if _, ok := appears[sub]; !ok {
            return false
        }

        if appears[sub] == words[sub] {
            // reaches maximum appears amount
            return false
        }

        appears[sub] += 1
    }
    return true
}

func findSubstring(s string, words []string) []int {
    ans := []int{}
    wordLen := len(words[0])
    concatenatedLen := wordLen * len(words)
    if len(s) < concatenatedLen {
        return ans
    }
    wordMap := map[string]int{}
    for _, w := range words {
        wordMap[w] += 1
    }

    // i, j := 0, 0
    for i := 0; i <= len(s) - concatenatedLen; i++ {
        if isStringConcatenated(s[i:i+concatenatedLen], wordLen, wordMap) {
            ans = append(ans, i)
            continue
        }
    }
    return ans
}

