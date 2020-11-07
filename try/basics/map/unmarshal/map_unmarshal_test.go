package unmarshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	err := Unmarshal()
	assert.NoError(t, err)
}
