package quiz

import (
	"errors"
	"math"
)

// Fermat's Last Theorem x**2 + y**2 = z**2
// 3**2 + 4**2 == 5**2 OR 6**2 + 8**2 == 10**2
// Input 10, 2 => [(3, 4, 5), (6, 8, 10)]
// (x or y <= 10)
// Input 10, 3 => []
// Input 10, 4 => []
// Input 10, 5 => []

var ErrOverSize = errors.New("over size error")

// x**3 + y**3 = z**3
func fermatLastTheorem(maxNum, squareNum int) ([][3]int, error) {
	result := make([][3]int, 0)

	for x := 1; x <= maxNum; x++ {
		for y := x + 1; y <= maxNum; y++ {
			powSum := math.Pow(float64(x), float64(squareNum)) + math.Pow(float64(y), float64(squareNum))
			z := math.Pow(powSum, 1.0/float64(squareNum))

			if z > math.MaxFloat64 {
				return nil, ErrOverSize
			}

			if !isInteger(z) {
				continue
			}

			if math.Pow(z, float64(squareNum)) == powSum {
				result = append(result, [3]int{x, y, int(z)})
			}
		}
	}

	return result, nil
}

func isInteger(n float64) bool {
	if n == float64(int64(n)) {
		return true
	} else {
		return false
	}
}
