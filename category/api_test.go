package category

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/uanderson/pockee/firebase"
	"testing"
)

func TestNewApi(t *testing.T) {
	e := echo.New()
	fb := &firebase.Firebase{}
	service := &Service{}

	api := NewApi(e, fb, service)

	assert.NotNil(t, api)
	assert.Equal(t, e, api.echo)
	assert.Equal(t, fb, api.firebase)
	assert.Equal(t, service, api.service)
}
