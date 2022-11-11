package coupang

import (
	"container/list"
	"strconv"
	"strings"
)

func Decode(encoded string) string {
	if len(encoded) >= 1 && strings.Contains(encoded, "[") {
		encodedSlice := encoded[0:]
		digitStack := list.New()
		nonDigitStack := list.New()

		currentIdx := 0

		for currentIdx < len(encoded) {
			if isDigit(encodedSlice[currentIdx]) {
				openBracketCharAt := strings.Index(encodedSlice[currentIdx:], "[")
				repeatedDigit, err := strconv.Atoi(encodedSlice[currentIdx : openBracketCharAt+currentIdx])

				if err != nil {
					panic("Wrong encoding format")
				}
				digitStack.PushBack(repeatedDigit)
				currentIdx = currentIdx + openBracketCharAt
			} else if isOpenBracket(encodedSlice[currentIdx]) || isAlphabet(encodedSlice[currentIdx]) {
				nonDigitStack.PushBack(encodedSlice[currentIdx : currentIdx+1])
				currentIdx += 1
			} else if isClosedBracket(encodedSlice[currentIdx]) {
				temporalString := ""
				for nonDigitStack.Back() != nil {
					topItemValue := nonDigitStack.Back().Value

					if topItemValue.(string) != "[" {
						temporalString = topItemValue.(string) + temporalString
						nonDigitStack.Remove(nonDigitStack.Back())
					} else {
						nonDigitStack.Remove(nonDigitStack.Back())
						break
					}
				}
				repeatedTimesEntry := digitStack.Back()
				repeatedTimes := repeatedTimesEntry.Value.(int)
				originalTemporalStr := temporalString
				for i := 0; i < repeatedTimes-1; i++ {
					temporalString = originalTemporalStr + temporalString
				}
				digitStack.Remove(digitStack.Back())
				nonDigitStack.PushBack(temporalString)
				currentIdx += 1
			}
		}

		finalDecodedString := ""
		for nonDigitStack.Front() != nil {
			finalDecodedString += nonDigitStack.Front().Value.(string)
			nonDigitStack.Remove(nonDigitStack.Front())
		}
		return finalDecodedString
	}
	return encoded
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlphabet(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isOpenBracket(ch byte) bool {
	return ch == '['
}

func isClosedBracket(ch byte) bool {
	return ch == ']'
}
