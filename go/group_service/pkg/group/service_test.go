package group

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessCode(t *testing.T) {
	code := NewAccessCode()
	assert.NotEmpty(t, code)
	assert.Len(t, code, 6)
	assert.Equal(t, code, strings.ToUpper(code))
}
