package autoid

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockReader struct{}

func (m *mockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("foo")
}

func TestNew(t *testing.T) {
	id := New()

	assert.Len(t, id, 20)
}

func TestNewPanic(t *testing.T) {
	oldReader := rand.Reader
	defer func() {
		rand.Reader = oldReader
	}()

	rand.Reader = &mockReader{}

	assert.PanicsWithValue(t, "could not generate the id: foo", func() {
		New()
	})
}
