package dxo

import (
	"encoding/pem"
	"testing"
)

func TestPemCodec(t *testing.T) {

	data1 := "foo is 福"
	data2 := "bar is 坝"

	pfile1 := new(PemFile)
	pfile2 := new(PemFile)

	b1 := new(pem.Block)
	b2 := new(pem.Block)

	b1.Type = "Foo"
	b1.Bytes = []byte(data1)

	b2.Type = "Bar"
	b2.Bytes = []byte(data2)

	pfile1.Add(b1)
	pfile1.Add(b2)

	txt, err := pfile1.Format()
	if err != nil {
		t.Error(err)
	}

	err = txt.Parse(pfile2)
	if err != nil {
		t.Error(err)
	}

	t.Logf("PemFileText = \n %v", txt.String())

	foo, err := pfile1.FindBlock("foo")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("foo =  [%v]", foo)
	}

}
