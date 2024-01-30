package smaz2

import (
	"encoding/base64"
	"testing"
)

func Test_Compress(t *testing.T) {
	compress := Compress("this world is beautiful")
	b64 := base64.StdEncoding.EncodeToString(compress)
	t.Log(b64)
}

func Test_Decompress(t *testing.T) {
	b, err := base64.StdEncoding.DecodeString("BwEHNZYBIM0BYdsCaWb7")
	if err != nil {
		t.Fatal(err)
	}
	decompress, err := Decompress(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(decompress)
}
