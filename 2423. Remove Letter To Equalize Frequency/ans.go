func equalFrequency(word string) bool {
    if len(word) == 2 {
        return true
    }

    cnt := map[byte]int{}

    mostTime := 0

    for _, c := range []byte(word) {
        cnt[c] += 1
        if cnt[c] > mostTime {
            mostTime = cnt[c]
        }
    }

    if len(cnt) == 1 {
        // only one letter
        return true
    }

    cntAmt := map[int]int{}
    for _, v := range cnt {
        cntAmt[v] += 1
    }

    if len(cntAmt) > 2 {
        // Only 2 valid cases.
        return false
    }

    // only two amount patterns

    if cntAmt[1] == 1 || cntAmt[1] == len(cnt) { // only one letter appears 1 time
        // Consider case like "aabbccd", or "aaaaaaab"

        return true
    }

    if cntAmt[mostTime] > 1 {
        return false
    }

    if cntAmt[mostTime - 1] != 0 {
        return true
    }

    return false

}
