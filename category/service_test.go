package category

import (
	"github.com/stretchr/testify/assert"
	"github.com/uanderson/pockee/database"
	"testing"
)

func TestNewService(t *testing.T) {
	db := &database.Database{}
	service := NewService(db)

	assert.NotNil(t, service)
	assert.NotNil(t, service.dao)
}
