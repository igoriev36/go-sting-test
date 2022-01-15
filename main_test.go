package main

import (
	"reflect"
	"strconv"
	"testing"
)

// estimate time: 10min, used time: 8min
func TestCheckValidity(t *testing.T) {
	input := "23-ab-4-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	result := testValidity(input)
	if !result {
		t.Errorf("Expected true, getting %q", strconv.FormatBool(result))
	}
	input = "23-a4b-4a-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	result = testValidity(input)
	if result {
		t.Errorf("Expected false, getting %q", strconv.FormatBool(result))
	}
}

// estimate time: 5min, used time: 5min
func TestAverageNumber(t *testing.T) {
	input := "23-ab-4-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	result := averageNumber(input)
	answer := int64(42)
	if result != answer {
		t.Errorf("Expected %d, getting %d", answer, result)
	}
}

// estimate time: 5min, used time: 5min
func TestWholeStory(t *testing.T) {
	input := "23-ab-4-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	result := wholeStory(input)
	answer := "ab cabaa haha howareyou trhowareyouthismorning"
	if result != answer {
		t.Errorf("Expected %q, getting %q", answer, result)
	}
}

// estimate time: 8min, used time: 7min
func TestStoryStats(t *testing.T) {
	input := "23-ab-4-cabaa-56-haha-75-howareyou-54-trhowareyouthismorning"
	shortest, longest, avgLen, avgWords := storyStats(input)
	if shortest != "ab" {
		t.Errorf("Expected ab, getting %q", shortest)
	}
	if longest != "trhowareyouthismorning" {
		t.Errorf("Expected trhowareyouthismorning, getting %q", longest)
	}
	if avgLen != 8 {
		t.Errorf("Expected 8, getting %d", avgLen)
	}
	if reflect.DeepEqual(avgWords, []string{"howareyou"}) {
		t.Errorf("Expected [howareyou], getting %q", avgWords)
	}
}

// estimate time: 5min, used time: 5min
func TestGenerage(t *testing.T) {
	input := generate(true)
	result := testValidity(input)
	if !result {
		t.Errorf("Expected true, getting %q", strconv.FormatBool(result))
	}
	input = generate(false)
	result = testValidity(input)
	if result {
		t.Errorf("Expected false, getting %q", strconv.FormatBool(result))
	}
}
