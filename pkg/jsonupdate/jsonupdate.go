package jsonupdate

import (
	"encoding/json"

	"github.com/tidwall/sjson"
)

type JsonUpdate struct {
	jsondata string
}

func (s *JsonUpdate) String() string {
	return s.jsondata
}

func (s *JsonUpdate) Pretty() string {
	var obj map[string]interface{}
	json.Unmarshal([]byte(s.jsondata), &obj)

	jsonData, _ := json.MarshalIndent(obj, "", "  ")

	return string(jsonData)
}

func (s *JsonUpdate) Set(key, value string) *JsonUpdate {
	s.jsondata, _ = sjson.Set(s.jsondata, key, value)
	return s
}

func (s *JsonUpdate) Deleete(path string) *JsonUpdate {
	s.jsondata, _ = sjson.Delete(s.jsondata, path)
	return s
}

func NewJsonUpdate(jsondata string) *JsonUpdate {
	return &JsonUpdate{
		jsondata: jsondata,
	}
}
