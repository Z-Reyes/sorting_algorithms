package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {

	fmt.Println("Performing Merge Sort")
	number := 10000000
	referenceList := make([]comparable, number)

	for t := 0; t < number; t++ {
		referenceList[t] = comparableInt(rand.Intn(number))
	}
	currentList := make([]comparable, number)
	for n := 0; n < b.N; n++ {
		copy(currentList, referenceList)
		currentList = mergeSort(currentList, number/2)
	}
}
func BenchmarkQuickSort(b *testing.B) {

	fmt.Println("Performing Quick Sort")
	number := 10000000
	referenceList := make([]comparable, number)

	for t := 0; t < number; t++ {
		referenceList[t] = comparableInt(rand.Intn(number))
	}
	currentList := make([]comparable, number)
	for n := 0; n < b.N; n++ {
		copy(currentList, referenceList)
		currentList = quickSort(currentList)
	}
}
func BenchmarkInsertionSort(b *testing.B) {
	fmt.Println("Performing Insertion Sort")
	number := 10000000
	referenceList := make([]comparable, number)

	for t := 0; t < number; t++ {
		referenceList[t] = comparableInt(rand.Intn(number))
	}
	currentList := make([]comparable, number)
	for n := 0; n < b.N; n++ {
		copy(currentList, referenceList)
		currentList = insertionSort(currentList)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	fmt.Println("Performing Selection Sort")
	number := 10000000
	referenceList := make([]comparable, number)

	for t := 0; t < number; t++ {
		referenceList[t] = comparableInt(rand.Intn(number))
	}
	currentList := make([]comparable, number)
	for n := 0; n < b.N; n++ {
		copy(currentList, referenceList)
		currentList = selectionSort(currentList)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	fmt.Println("Performing Bubble Sort")
	number := 10000000
	referenceList := make([]comparable, number)

	for t := 0; t < number; t++ {
		referenceList[t] = comparableInt(rand.Intn(number))
	}
	currentList := make([]comparable, number)
	for n := 0; n < b.N; n++ {
		copy(currentList, referenceList)
		currentList = bubbleSort(currentList)
	}
}
