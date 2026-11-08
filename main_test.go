package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "нулевая длина", size: 0},
		{name: "отрицательная длина", size: -10},
		{name: "один элемент", size: 1},
		{name: "несколько элементов", size: 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateRandomElements(test.size)

			if len(got) != test.size && test.size > 0 {
				t.Fatalf("generateRandomElements(%d) вернул слайс длиной %d, ожидалась длина %d", test.size, len(got), test.size)
			}

			if test.size <= 0 && len(got) != 0 {
				t.Fatalf("generateRandomElements(%d) должен вернуть пустой слайс, получено %d элементов", test.size, len(got))
			}

			for _, value := range got {
				if value <= 0 {
					t.Fatalf("сгенерировано неположительное число: %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{name: "пустой слайс", data: []int{}, expected: 0},
		{name: "nil слайс", data: nil, expected: 0},
		{name: "один элемент", data: []int{42}, expected: 42},
		{name: "максимум в начале", data: []int{100, 1, 2, 3}, expected: 100},
		{name: "максимум в конце", data: []int{1, 2, 3, 100}, expected: 100},
		{name: "максимум в середине", data: []int{1, 100, 2, 3}, expected: 100},
		{name: "одинаковые значения", data: []int{7, 7, 7}, expected: 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)

			if got != test.expected {
				t.Fatalf("maximum(%v) = %d, ожидалось %d", test.data, got, test.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{name: "пустой слайс", data: []int{}, expected: 0},
		{name: "nil слайс", data: nil, expected: 0},
		{name: "один элемент", data: []int{42}, expected: 42},
		{name: "меньше восьми элементов", data: []int{3, 9, 1, 7}, expected: 9},
		{name: "ровно восемь элементов", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 8},
		{name: "длина не делится на восемь", data: []int{10, 2, 3, 4, 5, 6, 7, 8, 9}, expected: 10},
		{name: "максимум в последнем срезе", data: []int{1, 2, 3, 4, 5, 6, 7, 100}, expected: 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)

			if got != test.expected {
				t.Fatalf("maxChunks(%v) = %d, ожидалось %d", test.data, got, test.expected)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := make([]int, 10000)

	for i := range data {
		data[i] = rand.Intn(1000000) + 1
	}

	want := maximum(data)
	got := maxChunks(data)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("maxChunks() = %d, maximum() = %d", got, want)
	}
}
