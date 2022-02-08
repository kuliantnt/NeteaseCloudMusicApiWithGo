package util

import (
	"fmt"
	"testing"
)

// Test_Weapi 测试weapi加密
//  @param t
func Test_Weapi(t *testing.T) {
	tests := []struct {
		name string
		data map[string]string
	}{}
	//赋值
	value := struct {
		name string
		data map[string]string
	}{
		name: "test",
		data: make(map[string]string),
	}
	value.data["test"] = "test2"
	tests = append(tests, value)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//加密
			weapitype := Weapi(tt.data)
			//打印
			fmt.Println(weapitype)
		})
	}
}
