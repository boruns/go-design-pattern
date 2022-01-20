package template

import (
	"testing"

	"github.com/boruns/go-design-pattern/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	tel := template.NewTelecomSms()
	err := tel.Send("test", 123456)
	assert.NoError(t, err)
}
