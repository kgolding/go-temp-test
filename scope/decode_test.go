package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	m, err := Decode([]byte("A1234567DHello\r"))
	assert.NoError(t, err)
	assert.Equal(t, PagerType('A'), m.Type)
	assert.True(t, m.Type.IsAlphanumeric())
	assert.Equal(t, 1234567, m.Capcode)
	assert.Equal(t, PagerBeep('D'), m.Beep)
	assert.True(t, m.Beep.IsValid())
	assert.Equal(t, "Hello", m.Message)

	m, err = Decode([]byte("N1234567C888\r"))
	assert.NoError(t, err)
	assert.Equal(t, PagerType('N'), m.Type)
	assert.True(t, m.Type.IsNumeric())
	assert.Equal(t, 1234567, m.Capcode)
	assert.Equal(t, PagerBeep('C'), m.Beep)
	assert.True(t, m.Beep.IsValid())
	assert.Equal(t, "888", m.Message)

	m, err = Decode([]byte(""))
	assert.Error(t, err)

	m, err = Decode([]byte("ZZZZZZZZZZZZZZZZZZZ"))
	assert.Error(t, err)
}
