// import "math"

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := map[string]bool{}

	for i := range banned {
		bannedMap[banned[i]] = true
	}

	splitters := map[byte]bool{}
	for _, c := range []byte("!?',;. ") {
		splitters[c] = true
	}

	words := map[string]int{}
	maxAppears := 0
	var ans string
	wordBuilder := strings.Builder{}

	idx := 0
	for idx <= len(paragraph) {
		if idx < len(paragraph) {
			c := paragraph[idx]
			if _, ok := splitters[c]; !ok {
				if c >= 'A' && c <= 'Z' {
					c += 'a' - 'A'
				}
                // fmt.Println("add char: ", string(c), string(paragraph[idx]))
				wordBuilder.WriteByte(c)
				idx++
				continue
			}
		}

		// a splitter appears
		word := wordBuilder.String()
		wordBuilder.Reset()
		if word != "" && bannedMap[word] == false {
			words[word] += 1
			if words[word] > maxAppears {
				maxAppears = words[word]
				ans = word
			}
		}
		idx++
	}

	return ans
}
