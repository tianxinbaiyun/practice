package consumer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleConsumer(t *testing.T) {
	for {
		err := SimpleConsumer()
		assert.NoError(t, err)
	}
}

func TestGetReader(t *testing.T) {
	for {
		err := GetReader()
		assert.NoError(t, err)
	}
}

func TestGroupConsumer(t *testing.T) {
	err := GroupConsumer()
	assert.NoError(t, err)
}
