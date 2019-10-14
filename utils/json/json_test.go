package jsontest

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

// 内置的json模块使用了反射tag的模式，效率较低
func TestEmbJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*e)

	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Error(err)
	}
	t.Log(e)

	if v, err := e.MarshalJSON(); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
}

func BenchmarkEmbJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
