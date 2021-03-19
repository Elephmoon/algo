package strings

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func IsUniqSymbols(str string) bool {
	if len(str) > 128 {
		return false
	}
	alphabet := make([]bool, 128)
	for _, symbol := range str {
		if alphabet[uint8(symbol)] {
			return false
		}
		alphabet[uint8(symbol)] = true
	}
	return true
}

func IsTransposition(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	if SortStr(str1) != SortStr(str2) {
		return false
	}
	return true
}

func SortStr(str string) string {
	splitWord := strings.Split(str, "")
	sort.Strings(splitWord)
	return strings.Join(splitWord, "")
}

func ReplaceSpace(str string) string {
	splitStr := strings.Split(str, "")
	for i, symbol := range splitStr {
		if symbol == " " {
			splitStr[i] = "%20"
		}
	}
	return strings.Join(splitStr, "")
}

func IsPalindrome(str string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	symbols := make(map[rune]uint8)
	str = reg.ReplaceAllString(str, "")
	for _, char := range str {
		symbols[char]++
	}
	var foundOdd bool
	for _, u := range symbols {
		if u%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

func CompressStr(str string) string {
	strBuilder := strings.Builder{}
	counter := 0
	for i := 0; i < len(str); i++ {
		counter++
		if i+1 >= len(str) || str[i] != str[i+1] {
			strBuilder.WriteString(string(str[i]))
			strBuilder.WriteString(strconv.Itoa(counter))
			counter = 0
		}
	}
	if strBuilder.Len() >= len(str) {
		return str
	}
	return strBuilder.String()
}
