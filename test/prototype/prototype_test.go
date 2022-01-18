package prototype

import (
	"testing"
	"time"

	"github.com/boruns/go-design-pattern/prototype"
	"github.com/stretchr/testify/assert"
)

func TestPrototypeClone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2022")
	testDTime := time.Now()
	words := prototype.Keywords{
		"testA": {
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": {
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": {
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
		"testD": {
			Word:      "testD",
			Visit:     4,
			UpdatedAt: &testDTime,
		},
	}
	now := time.Now()
	updateKeywords := []*prototype.Keyword{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
		{
			Word:      "testD",
			Visit:     4,
			UpdatedAt: &testDTime,
		},
	}
	got := words.Clone(updateKeywords)

	// 浅拷贝，指向的是同一个对象
	assert.Equal(t, got["testA"], words["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updateKeywords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
	assert.NotEqual(t, got["testD"], updateKeywords[1])
}
