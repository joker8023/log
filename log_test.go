package log

import "testing"

func TestLog(t *testing.T) {
	config := &Config{
		Project:     "test",
		Environment: "dev",
	}
	log := NewLog("123123", config)
	log.Printf("err:%s", "1231sadsad")

}
