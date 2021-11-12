package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDocumentCPF(t *testing.T) {
	var c = Customer{Document: 4417932834}
	x := c.IsDocumentCPF()
	assert.True(t, x)
}
