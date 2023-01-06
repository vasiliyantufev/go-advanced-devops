package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {

	assert.Equal(t, "test", "test",
		fmt.Sprintf("Incorrect location. Expect %s, got %s", "test", "test"))
}
