package prototype

import (
	"encoding/json"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	Word      string
	Visit     int
	UpdatedAt *time.Time
}

// Clone 使用序列化与反序列化进行拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		// 这里是浅拷贝  直接拷贝地址
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updateWords {
		newKeywords[word.Word] = word.Clone()
	}
	return newKeywords
}
