package compose

import (
	"testing"

	"github.com/boruns/go-design-pattern/compose"
	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	organization := compose.NewOrganization()
	got := organization.Count()
	assert.Equal(t, got, 20)
}
