package util

import (
	"testing"
)

// func TestB64Decode(t *testing.T) {
// 	decode := string(B64Decode("eyJhbGciOiJTZWNwMjU2azEiLCJ0eXAiOiJKV1QifQ=="))
// 	t.Log(decode)
// }

func TestConfig(t *testing.T) {
	config := Config()
	t.Log(config)
}

func TestUUID(t *testing.T) {
	result, err := UUID()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
