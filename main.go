package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10000000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс целых положительных чисел заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	numbers := make([]int, size)

	for i := range numbers {
		numbers[i] = rand.Intn(1000000) + 1
	}

	return numbers
}

// maximum находит максимальное значение в слайсе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks находит максимальное значение, разделяя слайс на CHUNKS частей
// и обрабатывая каждую часть в отдельной горутине.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var waitGroup sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		// Последний срез забирает остаток длины, чтобы не потерять элементы.
		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		chunk := data[startIndex:endIndex]
		waitGroup.Add(1)

		go func(chunk []int, index int) {
			defer waitGroup.Done()
			maxValues[index] = maximum(chunk)
		}(chunk, i)
	}

	waitGroup.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n\n", max, elapsed)

	fmt.Println("Ищем максимальное значение в 8 потоков")

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
