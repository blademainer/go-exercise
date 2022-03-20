package codec

import (
	jsonc "encoding/json"
	xmlc "encoding/xml"
)

//go:generate types-gen ./

// Coder types to generate components
type Coder string

type (
	json Coder // 将obj解析为json字符串
	xml  Coder // 将obj解析为xml字符串
)

func (j json) Marshal(v interface{}) (string, error) {
	marshal, err := jsonc.Marshal(v)
	return string(marshal), err
}

func (j json) Unmarshal(data string, v interface{}) error {
	return jsonc.Unmarshal([]byte(data), v)
}

func (x xml) Marshal(v interface{}) (string, error) {
	marshal, err := xmlc.Marshal(v)
	return string(marshal), err
}

func (x xml) Unmarshal(data string, v interface{}) error {
	return xmlc.Unmarshal([]byte(data), v)
}
