package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFailed(t *testing.T) {
	if Add(1,2) != 3 {
		t.Errorf("Result is incorrect")
	}
}

func TestAddSuccess(t *testing.T) {
	assert.Equal( t,Add(1,2), 3)
}