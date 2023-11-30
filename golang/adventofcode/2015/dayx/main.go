package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Get the file path from the command line arguments.
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: ./check_json_order <file_path>")
		return
	}
	file_path := args[0]

	// Read the JSON file into a slice of objects.
	//"/Users/vkad2506/Kaufland_Cloud_Config/content/config/dev1/dev1.json"
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}
	defer file.Close()

	var objects []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	err = json.NewDecoder(file).Decode(&objects)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the original array is sorted.
	isSorted := sort.SliceIsSorted(objects, func(i, j int) bool {
		return objects[i].Name < objects[j].Name
	})

	// If the original array is already sorted, then skip the for loop.
	if isSorted {
		fmt.Println("The objects are in alphabetical order. ✅︎")
		return
	} else {
		fmt.Println("The objects are not in alphabetical order. (x)")
	}

	// Sort the objects by the name field.
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].Name < objects[j].Name
	})

	// Check if the objects are in alphabetical order.
	isSorted = true
	for i := 1; i < len(objects); i++ {
		if objects[i-1].Name > objects[i].Name {
			isSorted = false
			break
		}
	}

	// Print the result.
	if isSorted {
		fmt.Println("The objects are in alphabetical order.")
	} else {
		fmt.Println("The objects are not in alphabetical order.")
	}
}
