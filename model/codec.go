package model

import (
	"encoding/json"
	"log"
)

func (s *Script) Encode() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res
}

func (s *Script) Decode(data []byte) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Script) String() string {
	return string(s.Encode())
}
