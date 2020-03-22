package leetcode

// 题目理解反了，是chars拼出words，不是words拼出chars
func countCharacters(words []string, chars string) int {
	charsMap := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		charsMap[chars[i]-'a']++
	}

	var ans int
	for i := 0; i < len(words); i++ {
		tmpCharMap := make([]int, len(charsMap))
		copy(tmpCharMap, charsMap)
		var contains = true
		for j := 0; j < len(words[i]); j++ {
			if tmpCharMap[words[i][j]-'a'] == 0 {
				contains = false
				break
			} else {
				tmpCharMap[words[i][j]-'a']--
			}
		}

		if contains {
			ans += len(words[i])
		}
	}

	return ans
}

func countCharacters1(words []string, chars string) int {
	wordsMap := make([][]int, len(words))
	for i := 0; i < len(words); i++ {
		wordMap := make([]int, 26)
		word := words[i]

		for j := 0; j < len(word); j++ {
			wordMap[word[j]-'a']++
		}

		wordsMap[i] = wordMap
	}

	charsMap := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		charsMap[chars[i]-'a']++
	}

	var ans = make(map[int]struct{})
	for i := 0; i < len(wordsMap); i++ {
		wordMap := wordsMap[i]
		for j := 0; j < len(charsMap); j++ {
			if charsMap[j] > 0 {
				if wordMap[j] > 0 {
					charsMap[j] = 0
					ans[i] = struct{}{}
				}
			}
		}
	}

	for i := 0; i < len(charsMap); i++ {
		if charsMap[i] != 0 {
			return 0
		}
	}

	var ret int
	for k := range ans {
		ret += len(words[k])
	}

	return ret
}
