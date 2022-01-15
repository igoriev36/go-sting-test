package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// testValidity is function to check the validity of given string accoridng to the given rule
// check will be done by regular expression
// estimate time: 15min, used time: 20min
func testValidity(value string) bool {
	sArr := strings.Split(value, "-")
	var r *regexp.Regexp
	for i, v := range sArr {
		if i%2 == 0 {
			r, _ = regexp.Compile(`^[0-9]*$`)
		} else {
			r, _ = regexp.Compile(`^[a-zA-Z]+$`)
		}
		if !r.Match([]byte(v)) {
			return false
		}
	}
	return true
}

// averageNumber is function to get average value of all numbers in the given string
// estimate time: 5min, used time: 3min
func averageNumber(value string) int64 {
	if !testValidity(value) {
		return 0
	}
	sArr := strings.Split(value, "-")
	total := int64(0)
	num := int64(0)
	for i, v := range sArr {
		if i%2 == 0 {
			vv, _ := strconv.ParseInt(v, 10, 64)
			total += vv
			num++
		}
	}
	return int64(math.Round(float64(total / num)))
}

// wholeStory is function to get all words in the given string and make a sentence by joining with space
// estimate time: 5min, used time: 3min
func wholeStory(value string) string {
	if !testValidity(value) {
		return ""
	}
	sArr := strings.Split(value, "-")
	storys := []string{}
	for i, v := range sArr {
		if i%2 == 1 {
			storys = append(storys, v)
		}
	}
	return strings.Join(storys, " ")
}

// storyStats is function to get the shortest and longest word, average length of all words and average length words in the given string
// estimate time: 10min, used time: 10min
func storyStats(value string) (string, string, int, []string) {
	if !testValidity(value) {
		return "", "", 0, []string{}
	}
	ws := wholeStory(value)
	words := strings.Split(ws, " ")
	longest := ""
	shortest := words[0]
	avgwords := []string{}
	totalLen := 0
	for _, v := range words {
		if len(v) > len(longest) {
			longest = v
		}
		if len(v) < len(shortest) {
			shortest = v
		}
		totalLen += len(v)
	}
	avgRound := int(math.Round(float64(totalLen) / float64(len(words))))
	for _, v := range words {
		if len(v) == avgRound {
			avgwords = append(avgwords, v)
		}
	}
	return shortest, longest, avgRound, avgwords
}

func generate(value bool) string {
	tl := rand.Intn(100) + 1
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	blocks := []string{}
	for i := 0; i < tl; i++ {
		blocks = append(blocks, strconv.Itoa(rand.Intn(10000)))
		word := make([]byte, rand.Intn(100)+1)
		for j := range word {
			word[j] = charset[rand.Intn(52)]
		}
		if !value {
			blocks = append(blocks, string(word)+"1")
		} else {
			blocks = append(blocks, string(word))
		}
	}
	return strings.Join(blocks, "-")
}

// main is bootstraping function
func main() {
	input := "23-ab-4-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	fmt.Println("result", testValidity(input))
	fmt.Println("2 result", averageNumber(input))
	fmt.Println("3 result", wholeStory(input))
	s, l, ar, aw := storyStats(input)
	fmt.Println("4 result", s, l, ar, aw)
}
