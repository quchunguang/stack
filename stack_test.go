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

func TestPop(t *testing.T) {
	s := New()
	s.Push(`test`)
	ss := s.Pop().(string)
	assert.Equal(t, `test`, ss)

	assert.True(t, s.Empty())
	nothing := s.Pop()
	assert.Nil(t, nothing)
}

func TestLen(t *testing.T) {
	s := New()
	assert.Equal(t, 0, s.Len())

	s.Push(`test`)
	assert.Equal(t, 1, s.Len())

	s.Pop()
	assert.Equal(t, 0, s.Len())
}

func TestContain(t *testing.T) {
	s := New()
	assert.False(t, s.Contain(1))

	s.Push(`test`)
	s.Push(1)
	assert.True(t, s.Contain(`test`))
	assert.False(t, s.Contain(`test1`))
	assert.False(t, s.Contain(nil))

	assert.True(t, s.Contain(1))
	assert.False(t, s.Contain(0))
}

func TestMap(t *testing.T) {
	s := New()

	s.Push(2)
	s.Push(1)

	var value int
	mapFunc := func(a interface{}) bool {
		if v, ok := a.(int); ok {
			return v == value
		}
		return false
	}

	value = 1
	b := s.Map(mapFunc)
	assert.Equal(t, 1, b.(int))

	value = 2
	c := s.Map(mapFunc)
	assert.Equal(t, 2, c.(int))

	value = 3
	d := s.Map(mapFunc)
	assert.Nil(t, d)
}
