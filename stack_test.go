package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	s := New()
	s.Push(`test`)
	s.Push(2)
	assert.Equal(t, 2, s.Len())
	i := s.Pop().(int)
	assert.Equal(t, 2, i)
	assert.Equal(t, 1, s.Len())
	ss := s.Pop().(string)
	assert.Equal(t, `test`, ss)
	assert.True(t, s.Empty())
}
