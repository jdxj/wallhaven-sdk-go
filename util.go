package wallhaven_sdk_go

import "unicode"

func TrimSpaceAndSplit(str string) []string {
	var (
		run    = []rune(str)
		result []string
		word   []rune
	)
	for i := 0; i <= len(run); i++ {
		if i == len(run) || unicode.IsSpace(run[i]) {
			if len(word) != 0 {
				result = append(result, string(word))
				word = word[:0]
			}
			continue
		}
		word = append(word, run[i])
	}
	return result
}
