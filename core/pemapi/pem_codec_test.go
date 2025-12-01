package pemapi

import (
	"net/http"
	"testing"
)

func TestRequestCodec(t *testing.T) {

	codec := new(CODEC)
	req1 := new(Request)
	query := make(map[string]string)
	props := make(map[string]string)
	env := make(map[string]string)

	// prepare

	query["x"] = "110"
	query["y"] = "25"
	query["z"] = "256"

	env["EN_VAR_1"] = "/this/is/a/path"
	env["EN_VAR_2"] = "1002233"
	env["EN_VAR_3"] = "ABCDEFG.0123456789_xyz"

	props["p.xxx.f1"] = "12345"
	props["p.xxx.f2"] = "/a/b/c/ddd"
	props["p.xxx.f3"] = "ABCDEFG.0123456789_xyz"

	req1.Method = http.MethodPost
	req1.User = "foo"
	req1.Host = "bar.example.com"
	req1.PathWant = "/Test/Request/Codec"
	req1.Query = query
	req1.Bytes = []byte("hello, this is TestRequestCodec")

	req1.SetContentType("text/demo")
	req1.Properties = props
	req1.Environment = env

	// encode

	t1, err := codec.EncodeRequest(req1)
	if err != nil {
		t.Error(err)
		return
	}

	// decode

	req2, err := codec.DecodeRequest(t1)
	if err != nil {
		t.Error(err)
		return
	}

	// encode (2)

	t2, err := codec.EncodeRequest(req2)
	if err != nil {
		t.Error(err)
		return
	}

	// log results
	t.Logf("text1 =\n %s", t1)
	t.Logf("text2 =\n %s", t2)
}

func TestResponseCodec(t *testing.T) {

}
