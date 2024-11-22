package main

import (
	"encoding/json"
	"fmt"
)

// Record構造体を定義
type Record struct {
	RequestID int `json:"request_id"`
	UserID    int `json:"user_id"`
	Priority  int `json:"priority"`
	RequestTS int `json:"request_ts"`
}

func main() {
	data := []byte(`{
		"request_id" : 1,
		"user_id" : 2,
		"priority" : 3,
		"request_ts" : 1709168560
	}`)

	var record Record
	err := json.Unmarshal(data, &record)
	if err != nil {
		fmt.Println("JSONの解析エラー:", err)
		return
	}

	fmt.Printf("解析されたRecord: %+v\n", record)
}