package rpc_test

import (
	"lspgo/rpc"
	"testing"
)

type EncodingExample struct{
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
	incoming := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	expected := 16
	actual, err := rpc.DecodeMessage([]byte(incoming))
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("Expected: %d Actual: %d", expected, actual)
	}
}