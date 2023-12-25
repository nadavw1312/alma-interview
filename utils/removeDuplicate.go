package utils

func RemoveDuplicates(ids []string) []string {
	// Create a map to store unique IDs
	uniqueMap := make(map[string]bool)

	// Create a new slice to store unique IDs
	var uniqueIDs []string

	// Iterate through the original slice
	for _, id := range ids {
		// Check if the ID is not already in the map
		if !uniqueMap[id] {
			// Add the ID to the map and the new slice
			uniqueMap[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	return uniqueIDs
}
