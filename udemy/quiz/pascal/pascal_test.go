package pascal

import (
	"testing"
)

func TestGeneratePascalTriangle(t *testing.T) {
	PrintPascalTriangle(GeneratePascalTriangle(8))
	// fmt.Println(center("10", 6, "*"))
	// fmt.Println(center("10", 7, "*"))
}
