package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name string
	Age string
}

func main()  {
	people := []Man{
		Man{"bob","23"},
		Man{"tom","24"},
		Man{"cindy","25"},
	}

	if data, err := json.Marshal(people); err == nil {
		fmt.Printf("%s\n", data)
	}

	m := map[string][]string{
		"level":   {"debug"},
		"message": {"File not found", "Stack overflow"},
	}

	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
	}

	// 可以添加json规则
	type DebugInfo struct {
		Level  string `json:"level,omitempty"` // Level解析为level,忽略空值
		Msg    string `json:"message"`         // Msg解析为message
		Author string `json:"-"`               // 忽略Author
	}


}