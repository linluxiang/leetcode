func countCharacters(words []string, chars string) int {
    charMap := [26]byte{}

    for _, c := range []byte(chars) {
        charMap[c - 'a'] += 1
    }

    ans := 0
    appears := make([]byte, 26)
    for _, word := range words {
        isGood := true
        clear(appears)
        for _, c := range []byte(word) {
            idx := c - 'a'
            if charMap[idx] == 0 {
                isGood = false
                break
            }
            if appears[idx] == charMap[idx] {
                isGood = false
                break
            }
            appears[idx] += 1
        }
        if isGood {
            ans += len(word)
        }
    }
    return ans
}
