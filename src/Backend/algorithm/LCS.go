package algorithm

import "github.com/shopspring/decimal"

func lcs(virus, human string) int {
	virLen := len(virus)
	humLen := len(human)
	var T = make([][]int, virLen+1)
	for i := range T {
		T[i] = make([]int, humLen+1)
	}

	for i := 0; i < virLen+1; i++ {
		for j := 0; j < humLen+1; j++ {
			if i == 0 || j == 0 {
				T[i][j] = 0
			} else if virus[i-1] == human[j-1] {
				T[i][j] = 1 + T[i-1][j-1]
			} else {
				T[i][j] = max(T[i-1][j], T[i][j-1])
			}
		}
	}
	return T[virLen][humLen]
}

func SimilarityLevel(virus, human string) decimal.Decimal {
	return decimal.NewFromFloat(float64((lcs(virus, human) * 100.0) / len(human)))
}
