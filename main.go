package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	str := "23-ab-48-caba-56-haha"
	fmt.Println(testValidity(str))
	fmt.Println(averageNumber(str))
	fmt.Println(wholeStory(str))
	fmt.Println(storyStats(str))
}

func testValidity(str string) bool {
	ss := strings.Split(str, "-")

	for i, s := range ss {

		if !isEven(i+1) && isDigit(s) {
			continue
		}

		if isEven(i+1) && !isDigit(s) {
			continue
		}

		return false
	}

	return true

}

func averageNumber(str string) float64 {
	sum := 0
	count := 0
	ss := strings.Split(str, "-")

	for i := 0; i < len(ss); i += 2 {
		sum += toInt(ss[i])
		count++
	}

	return float64(sum) / float64(count)
}

func wholeStory(str string) string {
	ss := strings.Split(str, "-")
	words := []string{}

	for i := 1; i < len(ss); i += 2 {
		word := ss[i]
		if !isDigit(word) {
			words = append(words, word)
		}
	}

	return strings.Join(words, " ")
}

func storyStats(str string) (string, string, float64, []string) {
	ss := strings.Split(str, "-")
	shortestWord := ss[1]
	longestWord := ss[1]
	sum := 0
	count := 0

	for i := 1; i < len(ss); i += 2 {
		word := ss[i]
		if len(word) < len(shortestWord) {
			shortestWord = word
		}

		if len(word) > len(longestWord) {
			longestWord = word
		}
		sum += len(word)
		count++
	}

	avg := float64(sum) / float64(count)

	avgRoundUp := math.Ceil(avg)
	avgRoundDown := math.Floor(avg)
	wordList := []string{}

	for i := 1; i < len(ss); i += 2 {
		word := ss[i]
		if len(word) == int(avgRoundUp) && len(word) == int(avgRoundDown) {
			wordList = append(wordList, word)
		}
	}

	return shortestWord, longestWord, avg, wordList
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}

	return false
}
