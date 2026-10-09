package msrpc

import (
	"os"
	"reflect"
	"testing"
)

func TestNetShareEnumAllResponseShareInfoList(t *testing.T) {
	data, err := os.ReadFile("testdata/netshare-enum-all-level1.bin")
	if err != nil {
		t.Fatal(err)
	}

	response := NetShareEnumAllResponseDecoder(data)

	if response.IsInvalid() {
		t.Fatal("captured response is invalid")
	}
	if response.IsIncomplete() {
		t.Fatal("captured response is incomplete")
	}

	got := response.ShareInfoList()

	want := []ShareInfo{
		{Name: "ADMIN$", Type: 0x80000000, Comment: "Удаленный Admin"},
		{Name: "C$", Type: 0x80000000, Comment: "Стандартный общий ресурс"},
		{Name: "FileZilla-3.69.5", Type: 0, Comment: ""},
		{Name: "IPC$", Type: 0x80000003, Comment: "Удаленный IPC"},
		{Name: "Users", Type: 0, Comment: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ShareInfoList() mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}
