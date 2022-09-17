package triangle

import (
	"fmt"
	"testing"
)

func TestGenerateRandTriangle(t *testing.T) {
	triangleList := GenerateRandTriangle(5, 20)
	Print(triangleList)
	fmt.Printf("minPath: %v\n", SumMinPath(triangleList))
}
