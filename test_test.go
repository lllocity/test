package main

import (
	"encoding/json"
	"testing"
)

type Record struct {
	RequestID int `json:"request_id"`
	UserID    int `json:"user_id"`
	Priority  int `json:"priority"`
	RequestTS int `json:"request_ts"`
}

func TestUnmarshalRecord(t *testing.T) {
	data := []byte(`{
		"request_id" : 1,
		"user_id" : 2,
		"priority" : 3,
		"request_ts" : 1709168560
	}`)

	var record Record
	err := json.Unmarshal(data, &record)
	if (err != nil) {
		t.Errorf("JSONの解析エラー: %v", err)
	}

	expected := Record{
		RequestID: 1,
		UserID:    2,
		Priority:  3,
		RequestTS: 1709168560,
	}

	if record != expected {
		t.Errorf("期待される結果: %+v, 実際の結果: %+v", expected, record)
	}
}