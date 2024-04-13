package heapsort

import (
	"fmt"
	"math"
	"strings"
)

func Print(heap heap) {
	var result strings.Builder
	level := 0
	maxLevel := int(math.Log2(float64(len(heap))))

	// Calculate the number of nodes and total spaces needed at the last level for full alignment
	nodesAtMaxLevel := int(math.Pow(2, float64(maxLevel)))
	totalSpaces := (nodesAtMaxLevel - 1) * 4 // 4 spaces between nodes at the last level

	for level <= maxLevel {
		numElements := int(math.Pow(2, float64(level)))
		startIndex := int(math.Pow(2, float64(level))) - 1
		endIndex := startIndex + numElements
		if endIndex > len(heap) {
			endIndex = len(heap)
		}

		// Calculate the spaces for alignment
		leadingSpaces := totalSpaces / int(math.Pow(2, float64(level+1)))             // Divide total spaces by 2^(level+1) for leading spaces
		betweenSpaces := totalSpaces/int(math.Pow(2, float64(level))) - leadingSpaces // Calculate between spaces

		// Build the string for the current level
		line := strings.Repeat(" ", leadingSpaces)
		for i := startIndex; i < endIndex; i++ {
			line += fmt.Sprintf("%d", heap[i])
			if i < endIndex-1 {
				line += strings.Repeat(" ", betweenSpaces)
			}
		}
		result.WriteString(line + "\n")

		// Prepare lines between current level nodes and next level
		if level < maxLevel {
			nextLevelCount := int(math.Pow(2, float64(level+1)))
			spaceBetweenConnectors := (betweenSpaces + 2*leadingSpaces) / nextLevelCount
			connectorLine := strings.Repeat(" ", leadingSpaces/2) // Start with half of the leading spaces

			for i := 0; i < numElements; i++ {
				// Add connectors for the current node to its children
				if 2*i+1 < nextLevelCount {
					connector := "/"
					if 2*i+2 < nextLevelCount {
						connector += strings.Repeat(" ", spaceBetweenConnectors-1) + "\\"
					}
					connectorLine += connector
					if i < numElements-1 {
						connectorLine += strings.Repeat(" ", betweenSpaces-1)
					}
				}
			}
			result.WriteString(connectorLine + "\n")
		}

		level++
	}
	fmt.Println(result.String())
}
