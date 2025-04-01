func romanToInt(s string) int {
    romeMap := map[string]int{
        "I": 1,
        "IV": 4,
        "V": 5,
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50,
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500,
        "CM": 900,
        "M": 1000,
    }
    specialMap := map[string]bool {
        "IV": true,
        "IX": true,
        "XL": true,
        "XC": true,
        "CD": true,
        "CM": true,
    }
    ans := 0
    for i:= 0; i < len(s); i++ {
        c := s[i]
        switch c {
            case: 'V', 'D', 'M':
                ans += romeMap[c]
            default:
                if i + 1 < len(s) {
                    test := string(c) + string(s[i+1])
                    if specialMap[test] == true {
                        ans += romeMap[test]
                        i += 1
                    } else {
                        ans += romeMap[c]
                    }
                } else {
                    ans += romeMap[c]
                }
        }
    }
    return ans
}
