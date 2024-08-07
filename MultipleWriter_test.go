package HyperUtility

import (
	"bytes"
	"testing"
)

func TestMultipleWriter_Write(t *testing.T) {
	one := bytes.NewBuffer(make([]byte, 0))
	two := bytes.NewBuffer(make([]byte, 0))

	mwriter := NewMultipleWriter()

	mwriter.AddWriter(one)
	mwriter.AddWriter(two)

	_, _ = mwriter.Write([]byte("abcdefg"))

	if one.String() != "abcdefg" {
		t.Fail()
	}

	if one.String() != two.String() {
		t.Fail()
	}
}
