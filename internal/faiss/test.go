package faiss

import (
	"fmt"

	"github.com/DataIntelligenceCrew/go-faiss"
)

func TestFAISS() {
	dim := 4
	index, err := faiss.NewIndexFlatL2(dim) // L2 distance index
	if err != nil {
		panic(err)
	}

	// Add two example vectors (flattened)
	vectors := []float32{
		1, 2, 3, 4, // first vector
		4, 3, 2, 1, // second vector
	}
	err = index.Add(vectors)
	if err != nil {
		panic(err)
	}

	// Search with query vector
	query := []float32{1, 2, 3, 4}
	distances, labels, err := index.Search(query, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Nearest neighbor ID:", labels[0])
	fmt.Println("Distance:", distances[0])
}
