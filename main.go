package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

func makeRandomSlice(numItems, max int) []Customer {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]Customer, numItems)

	for i := 0; i < numItems; i++ {
		id := fmt.Sprintf("C%d", i)
		numPurchases := random.Intn(max)
		slice[i] = Customer{id, numPurchases}
	}
	return slice
}

// Print at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []Customer) {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i].numPurchases > slice[i+1].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

func countingSort(A []Customer, max int) []Customer {
	var B = make([]Customer, len(A))
	var C = make([]int, max)

	// count each value
	for i := 0; i < len(A); i++ {
		C[A[i].numPurchases]++
	}

	// Convert the counts into the number of items less than or equal to each value.
	for j := 1; j < max; j++ {
		C[j] = C[j] + C[j-1]
	}

	// finally loop from back to front
	for k := len(A) - 1; k >= 0; k-- {
		numPurchases := A[k].numPurchases
		C[numPurchases]--
		B[C[numPurchases]] = A[k]
	}
	return B
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
