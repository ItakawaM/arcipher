package analyze

import (
	"math"
	"sort"
)

type factor struct {
	number int
	count  int
}

func processDistances(ngrams []nGramCount) []factor {
	distances := make([]int, 0)
	for _, ngram := range ngrams {
		for i := range len(ngram.positions) - 1 {
			distances = append(distances, ngram.positions[i+1]-ngram.positions[i])
		}
	}

	factors := make(map[int]int)
	for _, distance := range distances {
		for i := 2; i < int(math.Sqrt(float64(distance)))+1; i++ {
			if distance%i == 0 {
				factors[i] += 1
				if i*i != distance {
					factors[distance/i] += 1
				}
			}
		}
	}

	results := make([]factor, 0, len(factors))
	for key, value := range factors {
		results = append(results, factor{key, value})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})

	return results
}

// TODO: Implement key calculation
