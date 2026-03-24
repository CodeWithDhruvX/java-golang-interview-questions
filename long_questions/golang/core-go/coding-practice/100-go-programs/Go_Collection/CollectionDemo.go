package main

import "fmt"

/**
 * Demonstrates Go slice operations equivalent to Java Collection Interface.
 * 
 * Methods covered:
 * - append(), copy()
 * - manual remove, filter
 * - len(), cap()
 * - contains, containsAll
 * - to array conversion
 */

func main() {
    // Using slice as the Collection implementation
    collection := []string{}

    fmt.Println("--- Adding Elements ---")
    // append equivalent to add(E e)
    collection = append(collection, "Apple")
    collection = append(collection, "Banana")
    collection = append(collection, "Cherry")
    fmt.Printf("Collection after add: %v\n", collection)

    // append equivalent to addAll(Collection c)
    moreFruits := []string{"Date", "Elderberry"}
    collection = append(collection, moreFruits...)
    fmt.Printf("Collection after addAll: %v\n", collection)

    fmt.Println("\n--- Checking Status ---")
    // len() equivalent to size()
    fmt.Printf("Size: %d\n", len(collection))
    // isEmpty equivalent
    fmt.Printf("Is Empty: %t\n", len(collection) == 0)
    
    // contains equivalent
    fmt.Printf("Contains 'Banana': %t\n", contains(collection, "Banana"))
    fmt.Printf("Contains all 'moreFruits': %t\n", containsAll(collection, moreFruits))

    fmt.Println("\n--- Removing Elements ---")
    // remove equivalent (manual implementation)
    collection = remove(collection, "Apple")
    fmt.Printf("Removed 'Apple': %v\n", collection)

    // removeAll equivalent
    toRemove := []string{"Banana", "Date"}
    collection = removeAll(collection, toRemove)
    fmt.Printf("Removed multiple: %v\n", collection)

    fmt.Println("\n--- Other Operations ---")
    // clear equivalent
    collection = collection[:0] // Clear slice
    fmt.Printf("After clear: %v\n", collection)
    
    // toArray equivalent
    newCollection := []string{"X", "Y", "Z"}
    array := toArray(newCollection)
    fmt.Printf("To array: %v\n", array)
}

// Helper functions
func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func containsAll(slice, subSlice []string) bool {
    for _, item := range subSlice {
        if !contains(slice, item) {
            return false
        }
    }
    return true
}

func remove(slice []string, item string) []string {
    for i, s := range slice {
        if s == item {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func removeAll(slice, toRemove []string) []string {
    result := []string{}
    toRemoveSet := make(map[string]bool)
    for _, item := range toRemove {
        toRemoveSet[item] = true
    }
    
    for _, item := range slice {
        if !toRemoveSet[item] {
            result = append(result, item)
        }
    }
    return result
}

func toArray(slice []string) []string {
    // In Go, slices are already backed by arrays
    // This is just a copy demonstration
    result := make([]string, len(slice))
    copy(result, slice)
    return result
}
