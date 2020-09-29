package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Filter to get something in slice.
type Filter func([]int) int

// FilterRecursive to get something in slice.
type FilterRecursive func([]int, int, int) int

// Group something.
type Group interface {
	GetMax() int
	GetMin() int
	GetAverage() int
	GetTotal() int
	GetData() []int
}

// Collection of integer something.
type Collection struct {
	Max, Min, Average, Total int
	Data                     []int
}

// GetMax of Collection.
func (collection Collection) GetMax() int {
	return collection.Max
}

// GetMin of Collection.
func (collection Collection) GetMin() int {
	return collection.Min
}

// GetAverage of Collection.
func (collection Collection) GetAverage() int {
	return collection.Average
}

// GetTotal of Collection.
func (collection Collection) GetTotal() int {
	return collection.Total
}

// GetData of Collection
func (collection Collection) GetData() []int {
	return collection.Data
}

func filterMax(slice []int) int {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if max <= slice[i] {
			max = slice[i]
		}
	}
	return max
}

func filterMin(slice []int) int {
	min := slice[0]
	for i := 0; i < len(slice); i++ {
		if min >= slice[i] {
			min = slice[i]
		}
	}
	return min
}

func filterMinRecursive(slice []int, value, index int) int {
	if index == 0 {
		return value
	}
	if slice[index] <= value {
		value = slice[index]
	}
	return filterMinRecursive(slice, value, (index - 1))
}

func filterMaxRecursive(slice []int, value, index int) int {
	if index == 0 {
		return value
	}
	if slice[index] >= value {
		value = slice[index]
	}
	return filterMaxRecursive(slice, value, (index - 1))
}

func average(slice []int) int {
	total := sum(slice)
	return total / len(slice)
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func slicing(slices []int, amount int) [][]int {
	var chunks [][]int
	length := len(slices)
	divided := length / amount

	for i := 0; i < amount; i++ {
		iteration := i + 1
		finish := iteration * divided
		start := i * divided
		slice := slices[start:finish]
		chunks = append(chunks, slice)
	}
	return chunks
}

func getValue(slices []int, filter Filter) int {
	return filter(slices)
}

func getValueRecursive(slices []int, filter FilterRecursive) int {
	return filter(slices, slices[0], (len(slices) - 1))
}

func getMinTotal(slice []int) (int, int) {
	min := slice[0]
	key := 0
	for i := 0; i < len(slice); i++ {
		if min <= slice[i] {
			min = slice[i]
			key = i
		}
	}
	return key, min
}

func getMaxTotal(slice []int) (int, int) {
	max := slice[0]
	key := 0
	for i := 0; i < len(slice); i++ {
		if max >= slice[i] {
			max = slice[i]
			key = i
		}
	}
	return key, max
}

func divided(a, b int) (int, error) {
	if b == 0 {
		fmt.Println(a, "Dibagi", b)
		return 100, errors.New("Divided by zero")
	}
	fmt.Println(a, "Dibagi", b)
	result := a / b
	return result, nil
}

func assestSomething(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		random := rand.Intn(3)
		newValue, _ := divided(numbers[i], random)
		numbers[i] = newValue
	}
	return numbers
}

func printGroup(group Group, iteration int) {
	fmt.Println("Kumpulan ke-", iteration, group.GetData(), ". Rata-rata:", group.GetAverage(), ". Penjumlahan:", group.GetTotal(), ". Nilai Minimal:", group.GetMin(), ". Nilai Maksimal", group.GetTotal())
}

func main() {
	somethings := []int{23, 45, 67, 54, 66, 19, 56, 78, 89, 44, 11, 22, 33, 44, 55, 66, 77, 88, 99, 23, 34, 32, 23, 12}

	chunks := slicing(somethings, 3)
	totals := []int{
		0, 0, 0,
	}
	mins := [3]int{
		0, 0, 0,
	}
	maxs := [3]int{
		0, 0, 0,
	}
	averages := [3]int{
		0, 0, 0,
	}

	collections := [3]Collection{}

	for i := 0; i < len(chunks); i++ {
		iteration := i + 1
		chunk := chunks[i]
		averages[i] = average(chunk)
		totals[i] = sum(chunk)
		mins[i] = getValueRecursive(chunk, filterMinRecursive)
		maxs[i] = getValueRecursive(chunk, filterMaxRecursive)
		collections[i] = Collection{
			Max:     maxs[i],
			Min:     mins[i],
			Average: averages[i],
			Total:   totals[i],
			Data:    chunk,
		}
		printGroup(collections[i], iteration)
	}

	keyMin, minTotal := getMinTotal(totals)
	fmt.Println("Total Kumpulan terkecil adalah", minTotal, "oleh kumpulan", chunks[keyMin])

	keyMax, maxTotal := getMaxTotal(totals)
	fmt.Println("Total Kumpulan terbesar adalah", maxTotal, "oleh kumpulan", chunks[keyMax])

	for i := 0; i < len(chunks); i++ {
		fmt.Println("Kumpulan sebelum diolah", chunks[i])
		chunks[i] = assestSomething(chunks[i])
		fmt.Println("Kumpulan sesudah diolah", chunks[i])
	}

}
