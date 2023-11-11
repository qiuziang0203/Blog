package utils

import (
	"encoding/json"
)

type StringArray []string

// 序列化
func (m *StringArray) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *StringArray) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
