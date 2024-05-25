package rpc_test

import (
	"testing"

	"lspgo/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incoming := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	expectedMethod := "hi"
	expectedLength := 15
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if expectedLength != contentLength {
		t.Fatalf("Expected: %d Actual: %d", expectedLength, contentLength)
	}
	if expectedMethod != method {
		t.Fatalf("Expected: %s Actual: %s", expectedMethod, method)
	}
}
