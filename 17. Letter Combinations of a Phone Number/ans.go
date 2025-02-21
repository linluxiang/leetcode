func concat(digits string, prefix string, num2letter map[byte][]string, result *[]string) {
    if len(digits) == 0 {
        *result = append(*result, prefix)
        return
    }

    for _, letter := range num2letter[digits[0]] {
        concat(digits[1:len(digits)], prefix + letter, num2letter, result)
    }
}

func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    num2letter := map[byte][]string {
        '1': []string{},
        '2': []string{"a", "b", "c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m", "n", "o"},
        '7': []string{"p", "q", "r", "s"},
        '8': []string{"t", "u", "v"},
        '9': []string{"w", "x", "y", "z"},
    }

    result := []string{}

    concat(digits, "", num2letter, &result)
    return result
}
