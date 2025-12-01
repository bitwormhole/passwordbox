package pemapi

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type CODEC struct {
}

func (inst *CODEC) decodePemDoc(raw []byte) ([]*pem.Block, error) {

	list := make([]*pem.Block, 0)
	data2do := raw

	for {
		block, rest := pem.Decode(data2do)
		if block == nil {
			break
		}
		list = append(list, block)
		data2do = rest
	}

	return list, nil
}

func (inst *CODEC) encodePemDoc(doc []*pem.Block) ([]byte, error) {
	out := new(bytes.Buffer)
	for _, block := range doc {
		err := pem.Encode(out, block)
		if err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}

func (inst *CODEC) DecodeRequest(text RequestText) (*Request, error) {
	decoder := innerRequestDecoder{codec: inst}
	return decoder.decode(text)
}

func (inst *CODEC) DecodeResponse(text ResponseText) (*Response, error) {
	decoder := innerResponseDecoder{codec: inst}
	return decoder.decode(text)
}

func (inst *CODEC) DecodeMessage(text MessageText) (*Message, error) {
	decoder := innerMessageDecoder{codec: inst}
	return decoder.decode(text)
}

func (inst *CODEC) EncodeRequest(req *Request) (RequestText, error) {
	encoder := innerRequestEncoder{codec: inst}
	return encoder.encode(req)
}

func (inst *CODEC) EncodeResponse(resp *Response) (ResponseText, error) {
	encoder := innerResponseEncoder{codec: inst}
	return encoder.encode(resp)
}

func (inst *CODEC) EncodeMessage(msg *Message) (MessageText, error) {
	encoder := innerMessageEncoder{codec: inst}
	return encoder.encode(msg)
}

////////////////////////////////////////////////////////////////////////////////

type innerMessageEncoder struct {
	codec *CODEC
}

func (inst *innerMessageEncoder) encode(msg *Message) (MessageText, error) {

	list, err := inst.prepareBlockList(msg)
	if err != nil {
		return "", err
	}

	bin, err := inst.codec.encodePemDoc(list)
	if err != nil {
		return "", err
	}

	return MessageText(bin), nil
}

func (inst *innerMessageEncoder) prepareBlockList(msg *Message) ([]*pem.Block, error) {

	if msg == nil {
		return nil, fmt.Errorf("message is nil")
	}

	b0 := &msg.Block
	raw := msg.raw
	count := len(raw)

	if count > 0 {
		if b0 == raw[0] {
			return raw, nil
		}
	}

	dst := make([]*pem.Block, 0)
	dst = append(dst, b0)
	dst = append(dst, raw...)
	return dst, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerMessageDecoder struct {
	codec *CODEC
}

func (inst *innerMessageDecoder) decode(txt MessageText) (*Message, error) {

	bin := []byte(txt)
	blist, err := inst.codec.decodePemDoc(bin)
	if err != nil {
		return nil, err
	}

	msg := new(Message)
	msg.raw = blist
	count := len(blist)
	if count > 0 {
		b0 := blist[0]
		if b0 != nil {
			msg.Block = *b0
		}
	}
	return msg, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerRequestEncoder struct {
	codec *CODEC
}

func (inst *innerRequestEncoder) encode(req *Request) (RequestText, error) {

	if req == nil {
		return "", fmt.Errorf("request is nil")
	}

	method := inst.innerPrepareMethod(req)
	location := inst.innerPrepareLocation(req)
	clen := inst.innerPrepareContentLength(req)
	blockEnv := inst.innerPrepareEnvironmentBlock(req)
	blockProps := inst.innerPreparePropertiesBlock(req)
	msg := new(Message)

	*msg = req.Message

	msg.Type = BlockTypeRequest.String()
	msg.SetHeader("method", method)
	msg.SetHeader("location", location.String())
	msg.SetHeader("content-length", clen)

	if blockEnv != nil {
		msg.AddBlock(blockEnv)
	}
	if blockProps != nil {
		msg.AddBlock(blockProps)
	}

	txt, err := inst.codec.EncodeMessage(msg)
	return RequestText(txt), err
}

func (inst *innerRequestEncoder) innerPrepareMethod(req *Request) string {
	met := req.Method
	if met == "" {
		return http.MethodGet
	}
	return met
}

func (inst *innerRequestEncoder) innerPrepareLocation(req *Request) *url.URL {

	location := req.Location

	if location == nil {

		location = new(url.URL)
		path := req.PathWant
		user := url.User(req.User)
		query := inst.innerPrepareQueryString(req)
		scheme := req.Protocol

		if scheme == "" {
			scheme = "pemapi"
		}

		location.Scheme = scheme
		location.User = user
		location.Host = req.Host
		location.RawPath = path
		location.Path = path
		location.RawQuery = query
	}

	return location
}

func (inst *innerRequestEncoder) innerPrepareContentLength(req *Request) string {
	content := req.Bytes
	clen := len(content)
	return strconv.Itoa(clen)
}

func (inst *innerRequestEncoder) innerPrepareQueryString(req *Request) string {

	builder := new(strings.Builder)
	src := req.Query

	for name, value := range src {
		if name == "" || value == "" {
			continue
		}
		if builder.Len() > 0 {
			builder.WriteRune('&')
		}
		builder.WriteString(name)
		builder.WriteRune('=')
		builder.WriteString(value)
	}

	return builder.String()
}

func (inst *innerRequestEncoder) innerPreparePropertiesBlock(req *Request) *pem.Block {

	table := req.Properties
	if table == nil {
		return nil
	}

	block := new(pem.Block)
	block.Type = BlockTypeProperties.String()
	block.Headers = table
	return block
}

func (inst *innerRequestEncoder) innerPrepareEnvironmentBlock(req *Request) *pem.Block {

	table := req.Environment
	if table == nil {
		return nil
	}

	block := new(pem.Block)
	block.Type = BlockTypeEnvironment.String()
	block.Headers = table
	return block
}

////////////////////////////////////////////////////////////////////////////////

type innerRequestDecoder struct {
	codec *CODEC
}

func (inst *innerRequestDecoder) decode(txt RequestText) (*Request, error) {

	msg, err := inst.codec.DecodeMessage(MessageText(txt))
	if err != nil {
		return nil, err
	}

	location, err := inst.innerParseLocation(msg)
	if err != nil {
		return nil, err
	}

	method, err := inst.innerParseMethod(msg)
	if err != nil {
		return nil, err
	}

	query, err := inst.innerParseQuery(location)
	if err != nil {
		return nil, err
	}

	req := new(Request)
	req.Message = *msg

	req.Location = location
	req.Method = method

	req.Protocol = location.Scheme
	req.User = location.User.Username()
	req.Host = location.Hostname()
	req.PathHave = location.Path
	req.PathWant = location.Path
	req.Query = query

	env, err := inst.innerParseEnvironment(msg)
	if err == nil {
		req.Environment = env
	}

	props, err := inst.innerParseProperties(msg)
	if err == nil {
		req.Properties = props
	}

	return req, nil
}

func (inst *innerRequestDecoder) innerParseLocation(msg *Message) (*url.URL, error) {
	str := msg.GetHeader("location")
	if str == "" {
		return nil, fmt.Errorf("no Header named 'location'")
	}
	str = strings.TrimSpace(str)
	return url.Parse(str)
}

func (inst *innerRequestDecoder) innerParseMethod(msg *Message) (string, error) {
	str := msg.GetHeader("method")
	if str == "" {
		return "", fmt.Errorf("no Header named 'method'")
	}
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return str, nil
}

func (inst *innerRequestDecoder) innerParseContentLength(msg *Message) (int, error) {

	data := msg.Bytes
	len2 := len(data)
	str := msg.GetHeader("content-length")
	if str == "" {
		return len2, nil
	}

	len1, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("bad content-length: %s", err.Error())
	}

	if len1 != len2 {
		return 0, fmt.Errorf("bad content-length, have:%d want:%d", len2, len1)
	}

	return len2, nil
}

func (inst *innerRequestDecoder) innerParseContentType(msg *Message) (string, error) {
	value := msg.GetHeader("content-type")
	return value, nil
}

func (inst *innerRequestDecoder) innerParseQuery(location *url.URL) (map[string]string, error) {

	dst := make(map[string]string)
	return dst, nil

}

func (inst *innerRequestDecoder) innerParseProperties(msg *Message) (map[string]string, error) {

	dst := make(map[string]string)
	return dst, nil

}

func (inst *innerRequestDecoder) innerParseEnvironment(msg *Message) (map[string]string, error) {

	dst := make(map[string]string)
	return dst, nil

}

////////////////////////////////////////////////////////////////////////////////

type innerResponseEncoder struct {
	codec *CODEC
}

func (inst *innerResponseEncoder) encode(resp *Response) (ResponseText, error) {

	return "", fmt.Errorf("no impl")

}

////////////////////////////////////////////////////////////////////////////////

type innerResponseDecoder struct {
	codec *CODEC
}

func (inst *innerResponseDecoder) decode(txt ResponseText) (*Response, error) {

	return nil, fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////
// EOF
