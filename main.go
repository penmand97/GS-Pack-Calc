package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// Defined a RequestData struct to parse JSON requests
type RequestData struct {
	PackSizes []int `json:"packSizes"` // List of available pack sizes
	Items     int   `json:"items"`     // Number of items to pack
}

// ConsolidatePacks function consolidated packs to the most efficient configuration
func ConsolidatePacks(packSizes []int, packs map[int]int) map[int]int {
	// Created a new dictionary to store the consolidated pack counts
	newPacks := make(map[int]int)
	for _, size := range packSizes {
		newPacks[size] = 0 // Initialised all pack sizes with 0
	}

	// Looped through each pack size
	for i, size := range packSizes {
		if packs[size] == 0 {
			continue // Skipped if there were no packs of this size
		}
		if packs[size] > 1 {
			if i+1 < len(packSizes) {
				// If a larger pack size was available, incremented its count by 1
				newPacks[packSizes[i+1]] += 1
			} else {
				// Otherwise, retained the current pack size count
				newPacks[size] += packs[size]
			}
		} else if packs[size] == 1 {
			// If there was exactly one pack of this size, retained its count
			newPacks[size] += 1
		}
	}
	return newPacks // Returned the consolidated pack counts
}

// GetPacks function calculated the required number of each pack size for the given items
func GetPacks(packSizes []int, items int) map[int]int {
	// Created a reversed copy of the packSizes array
	reversedPackSizes := make([]int, len(packSizes))
	copy(reversedPackSizes, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(reversedPackSizes))) // Reversed the array

	// Initialised a dictionary with pack sizes and 0 counts
	packs := make(map[int]int)
	for _, size := range reversedPackSizes {
		packs[size] = 0
	}

	// Calculated the number of packs needed for each size
	for _, size := range reversedPackSizes {
		count := items / size // Calculated how many packs of this size were needed
		items = items % size  // Calculated the remaining items
		for i := 0; i < count; i++ {
			packs[size] += 1 // Incremented the count of packs of this size
		}
	}

	// If there were remaining items, added one of the smallest pack size
	if items > 0 {
		packs[reversedPackSizes[len(reversedPackSizes)-1]] += 1
	}

	// Sorted the pack sizes in ascending order
	var sortedPackSizes []int
	for size := range packs {
		sortedPackSizes = append(sortedPackSizes, size)
	}
	sort.Ints(sortedPackSizes) // Sorted pack sizes
	sortedPacks := make(map[int]int)
	for _, size := range sortedPackSizes {
		sortedPacks[size] = packs[size]
	}

	// Consolidated the packs to the most efficient configuration
	consolidatedPacks := ConsolidatePacks(packSizes, sortedPacks)
	fmt.Println(consolidatedPacks) // Printed the consolidated packs (for debugging)
	return consolidatedPacks       // Returned the consolidated packs
}

// consolidateHandler function handled incoming HTTP requests to consolidate packs
func consolidateHandler(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	// Parsed the JSON request body into the requestData struct
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Returned an error if the JSON was invalid
		return
	}

	// Calculated the consolidated packs
	consolidatedPacks := GetPacks(requestData.PackSizes, requestData.Items)

	// Set the response content type to JSON and encoded the response data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consolidatedPacks)
}

// main function set up the HTTP server and routes
func main() {
	// Handled requests to "/consolidate" with the consolidateHandler function
	http.HandleFunc("/consolidate", consolidateHandler)
	// Served static files (HTML, CSS) from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// Printed a message indicating that the server had started and was listening on port 8081
	fmt.Println("Server started at :8081")
	// Started the HTTP server on port 8081
	http.ListenAndServe(":8081", nil)
}
