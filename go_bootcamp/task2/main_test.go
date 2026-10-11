package main

import (
	"slices"
	"testing"
)

// обычное поведение с K меньшим, чем количество уникальных слов
func TestTopKFrequentWords_KLessThanUnique(t *testing.T) {
	got := TopKFrequentWords("apple banana apple orange banana apple grape", 2)
	want := []string{"apple", "banana"}

	if !slices.Equal(got, want) {
		t.Errorf("TopKFrequentWords(...) = %v; want %v", got, want)
	}
}

// передача пустого списка слов
func TestTopKFrequentWords_EmptyString(t *testing.T) {
	got := TopKFrequentWords("", 3)
	want := []string{}

	if !slices.Equal(got, want) {
		t.Errorf("TopKFrequentWords(...) = %v; want %v", got, want)
	}
}

// передача списка слов, где K больше, чем число уникальных слов
func TestTopKFrequentWords_KGreaterThanUnique(t *testing.T) {
	got := TopKFrequentWords("cat dog cat", 5)
	want := []string{"cat", "dog"}

	if !slices.Equal(got, want) {
		t.Errorf("TopKFrequentWords(...) = %v; want %v", got, want)
	}
}
