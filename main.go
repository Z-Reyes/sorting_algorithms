package main

import (
	"fmt"
	"math/rand"
)

type comparable interface {
	getValue() int
}

type comparableInt int

func (c comparableInt) getValue() int {
	return int(c)
}

func Print(c []comparable) {
	fmt.Printf("\n[")
	for _, val := range c {
		fmt.Print(val, ", ")
	}
	fmt.Println("]")
}

func isSortedAscending(c []comparable) bool {
	idxStart := 0
	idxEnd := len(c) - 1
	for idxStart < idxEnd {

		if c[idxStart].getValue() > c[idxStart+1].getValue() {
			return false
		}
		idxStart++
	}
	return true
}

//Performs bubble sort on array of comparables. Will take a long time. Produces an ascending list.
func bubbleSort(c []comparable) []comparable {
	isSorted := false
	endIndex := len(c) - 1

	for !isSorted {
		swapHappened := false
		idx := 0
		for idx < endIndex {
			if c[idx].getValue() > c[idx+1].getValue() {
				temp := c[idx]
				c[idx] = c[idx+1]
				c[idx+1] = temp
				swapHappened = true
			}
			idx++
		}
		if !swapHappened {
			isSorted = true
		}
		endIndex--

	}
	return c
}

func selectionSort(c []comparable) []comparable {

	endIndex := len(c)
	for minIdx, minVal := range c {
		curMinIndex := minIdx
		curMinVal := minVal.getValue()
		curIndex := curMinIndex + 1 //offset by one so we don't doublecheck
		for curIndex < endIndex {
			if c[curIndex].getValue() < curMinVal {
				curMinVal = c[curIndex].getValue()
				curMinIndex = curIndex
			}
			curIndex++
		}
		c[minIdx] = c[curMinIndex]
		c[curMinIndex] = minVal
	}
	return c
}

func insertionSort(c []comparable) []comparable {
	currentHigh := 0
	for i, val := range c {
		//fmt.Println(currentHigh)
		if i == 0 {
			continue
		}
		if val.getValue() < c[currentHigh].getValue() {
			//Then do the reverse order insertion plus shifting.
			//First, find the correct index to insert val.
			startingPos := currentHigh
			for ; startingPos >= 0; startingPos-- {
				if startingPos == 0 {
					break
				}
				if c[startingPos].getValue() < val.getValue() {
					startingPos++
					break
				}
			}
			temp := val
			for pointy := currentHigh + 1; pointy > startingPos; pointy-- {
				c[pointy] = c[pointy-1]
			}
			c[startingPos] = temp

		}
		currentHigh++
	}
	return c
}

func mergeSort(c []comparable, parallelLimit int) []comparable {
	var midPoint int = int(len(c)) / 2
	if len(c) == 2 {
		if c[0].getValue() > c[1].getValue() {
			temp := c[0]
			c[0] = c[1]
			c[1] = temp
		}
		return c
	} else if midPoint >= 1 {
		leftArray := make([]comparable, len(c[:midPoint]))
		rightArray := make([]comparable, len(c[midPoint:]))
		if len(c) > parallelLimit {
			leftChannel := make(chan []comparable, len(c[:midPoint]))
			rightChannel := make(chan []comparable, len(c[midPoint:]))
			go mergeSortParallel(c[:midPoint], parallelLimit, leftChannel)
			go mergeSortParallel(c[midPoint:], parallelLimit, rightChannel)
			//Now, we need to wait until both subroutines finish.
			copy(leftArray, <-leftChannel)
			copy(rightArray, <-rightChannel)
		} else {
			leftArray = mergeSort(c[:midPoint], parallelLimit)
			rightArray = mergeSort(c[midPoint:], parallelLimit)
		}

		//Now, we merge the two arrays together
		leftDeepCopy := make([]comparable, len(leftArray))
		copy(leftDeepCopy, leftArray)

		rightDeepCopy := make([]comparable, len(rightArray))
		copy(rightDeepCopy, rightArray)

		leftArrayPointer := 0
		rightArrayPointer := 0
		cPointer := 0
		for ; cPointer < len(c); cPointer++ {
			if leftArrayPointer == len(leftDeepCopy) || rightArrayPointer == len(rightDeepCopy) {
				//If one pointer is at the end, then it signifies we'll need to at least wholesale copy the rest of one array to c.
				for ; rightArrayPointer < len(rightDeepCopy); rightArrayPointer, cPointer = rightArrayPointer+1, cPointer+1 {
					c[cPointer] = rightDeepCopy[rightArrayPointer]
				}
				for ; leftArrayPointer < len(leftDeepCopy); leftArrayPointer, cPointer = leftArrayPointer+1, cPointer+1 {
					c[cPointer] = leftDeepCopy[leftArrayPointer]
				}
				return c
			}

			if leftDeepCopy[leftArrayPointer].getValue() < rightDeepCopy[rightArrayPointer].getValue() {
				c[cPointer] = leftDeepCopy[leftArrayPointer]
				leftArrayPointer++
			} else {
				c[cPointer] = rightDeepCopy[rightArrayPointer]
				rightArrayPointer++
			}

		}

	}
	return c
}

//Basically the same as mergeSort, except that it returns to a channel as opposed to mergeSort.
func mergeSortParallel(c []comparable, parallelLimit int, output chan []comparable) {
	var midPoint int = int(len(c)) / 2
	if len(c) == 2 {
		if c[0].getValue() > c[1].getValue() {
			temp := c[0]
			c[0] = c[1]
			c[1] = temp
		}
		output <- c
		return
	} else if midPoint >= 1 {
		leftArray := make([]comparable, len(c[:midPoint]))
		rightArray := make([]comparable, len(c[midPoint:]))
		if len(c) > parallelLimit {
			leftChannel := make(chan []comparable)
			rightChannel := make(chan []comparable)

			go mergeSortParallel(c[:midPoint], parallelLimit, leftChannel)

			go mergeSortParallel(c[midPoint:], parallelLimit, rightChannel)
			//Now, we need to wait until both subroutines finish.
			copy(leftArray, <-leftChannel)
			copy(rightArray, <-rightChannel)
		} else {
			leftArray = mergeSort(c[:midPoint], parallelLimit)
			rightArray = mergeSort(c[midPoint:], parallelLimit)

		}
		//Now, we merge the two arrays together
		leftDeepCopy := make([]comparable, len(leftArray))
		copy(leftDeepCopy, leftArray)

		rightDeepCopy := make([]comparable, len(rightArray))
		copy(rightDeepCopy, rightArray)
		leftArrayPointer := 0
		rightArrayPointer := 0
		cPointer := 0
		for ; cPointer < len(c); cPointer++ {
			if leftArrayPointer == len(leftDeepCopy) || rightArrayPointer == len(rightDeepCopy) {
				//If one pointer is at the end, then it signifies we'll need to at least wholesale copy the rest of one array to c.
				for ; rightArrayPointer < len(rightDeepCopy); rightArrayPointer, cPointer = rightArrayPointer+1, cPointer+1 {
					c[cPointer] = rightDeepCopy[rightArrayPointer]
				}
				for ; leftArrayPointer < len(leftDeepCopy); leftArrayPointer, cPointer = leftArrayPointer+1, cPointer+1 {
					c[cPointer] = leftDeepCopy[leftArrayPointer]
				}
				output <- c
				return
			}

			if leftDeepCopy[leftArrayPointer].getValue() < rightDeepCopy[rightArrayPointer].getValue() {
				c[cPointer] = leftDeepCopy[leftArrayPointer]
				leftArrayPointer++
			} else {
				c[cPointer] = rightDeepCopy[rightArrayPointer]
				rightArrayPointer++
			}

		}

	}
	output <- c
}

func quickSort(c []comparable) []comparable {

	if len(c) < 2 {
		return c
	}
	pivot := rand.Intn(len(c) - 1)
	lowerPoint := 0
	initialTemp := c[pivot]
	c[pivot] = c[len(c)-1]
	c[len(c)-1] = initialTemp
	//Idea. Move pivot to the end of the slice before scanning. Then we don't need to worry about it colliding with the lowerPoint location.
	for scanner := 0; scanner < len(c)-1; scanner++ {
		if c[scanner].getValue() < c[len(c)-1].getValue() {
			//swap the values at scanner and lowerPoint. then increment lowerPoint
			temp := c[lowerPoint]
			c[lowerPoint] = c[scanner]
			c[scanner] = temp
			lowerPoint++
		}
	}
	initialTemp = c[len(c)-1]
	c[len(c)-1] = c[lowerPoint]
	c[lowerPoint] = initialTemp
	quickSort(c[:lowerPoint])
	if lowerPoint+1 < len(c) {
		quickSort(c[lowerPoint+1:])
	}
	return c
}

func main() {
	c := []comparable{comparableInt(5), comparableInt(2), comparableInt(1), comparableInt(4), comparableInt(3)}
	d := []comparable{comparableInt(5), comparableInt(2), comparableInt(1), comparableInt(4), comparableInt(3)}
	e := []comparable{comparableInt(5), comparableInt(2), comparableInt(1), comparableInt(4), comparableInt(3)}
	f := []comparable{comparableInt(5), comparableInt(2), comparableInt(1), comparableInt(3), comparableInt(4), comparableInt(6), comparableInt(10), comparableInt(9), comparableInt(7), comparableInt(8)}
	g := []comparable{comparableInt(5), comparableInt(2), comparableInt(10), comparableInt(3), comparableInt(4), comparableInt(6), comparableInt(1), comparableInt(9), comparableInt(7), comparableInt(8)}
	fmt.Println("C bubble sort")
	c = bubbleSort(c)
	fmt.Println("D selection sort")
	d = selectionSort(d)
	fmt.Println("E insertion sort")
	e = insertionSort(e)
	fmt.Println("F merge sort")
	f = mergeSort(f, 5)
	fmt.Println("G quick sort")
	g = quickSort(g)
	fmt.Println(isSortedAscending(c))
	fmt.Println(isSortedAscending(d))
	fmt.Println(isSortedAscending(e))
	fmt.Println(isSortedAscending(f))
	fmt.Println(isSortedAscending(g))

}
