package models

// unmarshal rueidis.RedisMessage for idx
type SearchIndex struct {
	Value     string `json:"Value"`
	IndexType string `json:"Type"`
}