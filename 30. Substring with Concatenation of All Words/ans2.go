import "fmt"

func findSubstring(s string, words []string) []int {

	// Check the explanation here: https://leetcode.com/problems/substring-with-concatenation-of-all-words/solutions/1753357/Clear-solution!-Easy-to-understand-with-diagrams//
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

    // ansMap := map[int]bool{}

	// 这里i从0到wordLen-1的原因是因为，因为word的长度固定，所以在滑动窗口从i开始移动的过程中，因为每次移动一个wordLen，所以i+wordLen已经判断过了
	for i := 0; i < wordLen; i++ {
		appears := map[string]int{}
		for _, k := range words {
			appears[k] = 0
		}

		left, right := i, i-1

		for right < len(s) {
			// Right 往右扩张
			right += wordLen
            if right >= len(s) {
                break
            }
			sub := s[right-wordLen + 1 : right + 1]
            // fmt.Println("right: ", right, " left: ", left, " sub: ", sub, " appears: ", appears, " wordMap: ", wordMap)

			if _, ok := appears[sub]; !ok {
				// this word doesn't exist
				left = right + 1
				right = right
                for _, k := range words {
			        appears[k] = 0
		        }
				continue

			}

            appears[sub] += 1

			for appears[sub] > wordMap[sub] {
				//某个word出现次数过多
                leftSub := s[left:left+wordLen]
                appears[leftSub] -= 1
				left += wordLen
			}


			if right-left+1 == concatenatedLen {
                equal := true
                for k := range wordMap {
                    if wordMap[k] != appears[k] {
                        equal=false
                        break
                    }
                }
                if equal {
                    ans = append(ans, left)

                }
                sub := s[left:left+wordLen]
                appears[sub] -= 1
                left += wordLen
			}
		}
	}
	return ans
}
