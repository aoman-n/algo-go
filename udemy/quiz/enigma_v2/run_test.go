package enigmaV2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	err := Run()
	assert.NoError(t, err)
}
