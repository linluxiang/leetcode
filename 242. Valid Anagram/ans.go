func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    appears := map[rune]int{}
    for _, r := range s {
        appears[r] += 1
    }

    appears2 := map[rune]int{}

    for _, r := range t {
        appears2[r] += 1
        if _, ok := appears[r]; !ok {
            return false
        } 
        
        if appears[r] < appears2[r] {
            return false
        }
    }
    
    if len(appears) != len(appears2) {
        return false
    }

    for k := range appears {
        if appears[k] != appears2[k] {
            return false
        }
    }

    return true
}
