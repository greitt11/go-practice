package main

import "errors"

func sumEven(nums []int) int {
	answer := 0
	for _, num := range nums {
		if num%2 == 0 {
			answer += num
		}
	}
	return answer
}

func countWords(words []string) map[string]int {
	result := make(map[string]int)
	for _, word := range words {
		result[word]++
	}
	return result
}

func average(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("пустой слайс")
	}
	total := 0

	for _, num := range nums {
		total += num
	}
	result := float64(total) / float64(len(nums))

	return result, nil
}

func minMax(nums []int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("пустой слайс")
	}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min, max, nil
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		nowtarget := target - num
		v, ok := m[nowtarget]
		if ok {
			return []int{v, i}
		}
		m[num] = i

	}
	return nil
}

type Book struct {
	Title  string
	Author string
	Year   int
	Pages  int
}

func oldestBook(books []Book) (Book, error) {
	if len(books) == 0 {
		return Book{}, errors.New("пустой слайс")
	}
	oldest_i, oldest := 0, books[0].Year
	for i, book := range books {
		if book.Year < oldest {
			oldest = book.Year
			oldest_i = i
		}
	}
	return books[oldest_i], nil
}

func totalPages(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.Pages
	}
	return total
}

func booksByAuthor(books []Book) map[string]int {
	m := make(map[string]int)
	for _, book := range books {
		m[book.Author]++
	}
	return m
}

func main() {}
