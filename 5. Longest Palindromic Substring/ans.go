func expand(s string, i, j int) string {
	for i >= 0 && j <= len(s)-1 {
		if s[i] != s[j] {
			return s[i+1 : j]
		}
        i--
		j++
	}
    return s[i+1:j]
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
    result := ""
    for i := 0; i < len(s); i++ {
        r := expand(s, i, i)
        if len(r) > len(result) {
            result = r
        } 
    }

    for i:=0; i < len(s) - 1; i++ {
        if s[i] == s[i+1] {
            r := expand(s, i, i+1)
            if len(r) > len(result) {
                result = r
            }
        }
    }

    return result
}
