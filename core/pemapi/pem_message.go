package pemapi

import (
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/bitwormhole/passwordbox/core/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type MessageText string

type RequestText MessageText

type ResponseText MessageText

////////////////////////////////////////////////////////////////////////////////

func MakeMessageText(str string) MessageText {
	return MessageText(str)
}

func MakeRequestText(str string) RequestText {
	return RequestText(str)
}

func MakeResponseText(str string) ResponseText {
	return ResponseText(str)
}

func (txt MessageText) String() string {
	return string(txt)
}

func (txt RequestText) String() string {
	return string(txt)
}

func (txt ResponseText) String() string {
	return string(txt)
}

////////////////////////////////////////////////////////////////////////////////

// Message 表示 一条 pemapi 消息, 它的格式类似于一个完整的 PEM 文件, 由若干个 pem.block 组成
type Message struct {

	// data of block[0]
	pem.Block

	// list all blocks
	raw []*pem.Block
}

func (inst *Message) innerGetHeaders(create bool) map[string]string {
	table := inst.Headers
	if table == nil {
		if create {
			table = make(map[string]string)
			inst.Headers = table
		}
	}
	return table
}

func (inst *Message) innerGetBlockID(b *pem.Block) string {
	if b == nil {
		return ""
	}
	headers := b.Headers
	if headers == nil {
		return ""
	}
	return headers["id"]
}

func (inst *Message) GetHeader(name string) string {
	t := inst.innerGetHeaders(false)
	if t == nil {
		return ""
	}
	return t[name]
}

func (inst *Message) SetHeader(name, value string) {

	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	value = strings.TrimSpace(value)

	if name == "" {
		return
	}

	t := inst.innerGetHeaders(true)
	t[name] = value
}

func (inst *Message) SetContentType(ctype string) {
	inst.SetHeader("content-type", ctype)
}

func (inst *Message) SetContentLength(length int) {
	value := strconv.Itoa(length)
	inst.SetHeader("content-length", value)
}

// 使用 'CODEC' 代替:

// func (inst *Message) Encode() ([]byte, error) {
// 	all := inst.raw
// 	codec := new(CODEC)
// 	return codec.encodePemDoc(all)
// }

// func (inst *Message) Decode(data []byte) error {
// 	codec := new(CODEC)
// 	blist, err := codec.decodePemDoc(data)
// 	if err != nil {
// 		return err
// 	}
// 	inst.raw = blist
// 	return nil
// }

// func (inst *Message) EncodeString() (string, error) {
// 	bin, err := inst.Encode()
// 	if err != nil {
// 		return "", err
// 	}
// 	str := string(bin)
// 	return str, nil
// }

// func (inst *Message) DecodeString(data string) error {
// 	bin := []byte(data)
// 	return inst.Decode(bin)
// }

func (inst *Message) AddBlock(b *pem.Block) {
	if b == nil {
		return
	}
	list := inst.raw
	list = append(list, b)
	inst.raw = list
}

// 排除重复的 block, by [block.headers.id], 保留最后一个
func (inst *Message) RemoveRepeatedBlocks() {

	src := inst.raw
	tmp := make(map[string]*pem.Block)
	dst := make([]*pem.Block, 0)

	for _, block := range src {
		id := inst.innerGetBlockID(block)
		if id == "" {
			dst = append(dst, block)
		} else {
			tmp[id] = block
		}
	}

	for _, block := range tmp {
		dst = append(dst, block)
	}

	inst.raw = dst
}

func (inst *Message) ListBlocks() []*pem.Block {
	return inst.raw
}

func (inst *Message) Reset() {
	b0 := new(pem.Block)
	inst.Block = *b0
	inst.raw = nil
}

////////////////////////////////////////////////////////////////////////////////

type Status struct {
	Code int
	Text string
}

func (inst *Status) String() string {
	code := inst.Code
	str1 := strconv.Itoa(code)
	str2 := http.StatusText(code)
	return str1 + " " + str2
}

func ParseStatus(str string) (*Status, error) {

	idx := strings.IndexByte(str, ' ')
	if idx < 1 {
		return nil, fmt.Errorf("bad status string: [%s]", str)
	}

	str1 := strings.TrimSpace(str[0:idx])
	str2 := strings.TrimSpace(str[idx+1:])
	code, err := strconv.Atoi(str1)
	if err != nil {
		return nil, fmt.Errorf("bad status code: %s", err.Error())
	}

	sta := new(Status)
	sta.Code = code
	sta.Text = str2
	return sta, nil
}

////////////////////////////////////////////////////////////////////////////////

type Request struct {

	// Message.Type = "PEMAPI REQUEST"
	Message

	Method   string
	Location *url.URL

	PathHave string // 匹配所得的路径
	PathWant string // = Location.Path
	User     string // = Location.User
	Host     string // = Location.Host
	Protocol string // = Location.Scheme

	Query       map[string]string // = Location.Query
	Params      map[string]string // = Location.Path_params
	Environment map[string]string // @block(Environment)
	Properties  map[string]string // @block(Properties)

	// user@host = 目标 bank 的名称
	Bank dxo.EmailAddress
}

func (inst *Request) SetLocation(location string) error {
	u, err := url.Parse(location)
	if err != nil {
		return err
	}
	inst.Location = u
	return nil
}

// [弃用]

// func (inst *Request) Text() RequestText {

// 	b1 := inst.innerMakeRequestBlock()
// 	inst.AddBlock(b1)
// 	inst.RemoveRepeatedBlocks()

// 	str, err := inst.Message.EncodeString()
// 	if err != nil {
// 		str = err.Error()
// 	}
// 	return RequestText(str)
// }

// [弃用]
// func (inst *Request) innerMakeRawQuery(src map[string]string, dst *url.URL) {

// 	sep := ""
// 	builder := new(strings.Builder)

// 	for k, v := range src {
// 		builder.WriteString(sep)
// 		builder.WriteString(k)
// 		builder.WriteString("=")
// 		builder.WriteString(v)
// 		sep = "&"
// 	}

// 	dst.RawQuery = builder.String()
// }

// [弃用]
// func (inst *Request) innerMakeRequestBlock() *pem.Block {

// 	query := inst.Query
// 	params := inst.Params
// 	body := inst.Bytes
// 	u1 := new(url.URL)
// 	u2 := new(url.URL) // for params

// 	u1.Path = inst.PathWant
// 	inst.innerMakeRawQuery(query, u1)
// 	inst.innerMakeRawQuery(params, u2)

// 	inst.SetHeader("method", inst.Method)
// 	inst.SetHeader("url", u1.String())
// 	inst.SetHeader("params", u2.String())
// 	inst.SetHeader("id", "request")
// 	inst.SetContentLength(len(body))

// 	inst.Block.Type = BlockTypeRequest.String()
// 	block := new(pem.Block)
// 	*block = inst.Block

// 	return block
// }

// [弃用]
// func (inst *Request) innerGetQueryMap(u *url.URL) map[string]string {
// 	src := u.Query()
// 	dst := make(map[string]string)
// 	for name, values := range src {
// 		for _, val := range values {
// 			dst[name] = val
// 		}
// 	}
// 	return dst
// }

// [弃用]
// func (inst *Request) Parse(txt RequestText) error {

// 	const wantType = BlockTypeRequest

// 	str := txt.String()
// 	bin := []byte(str)
// 	err := inst.Message.Decode(bin)
// 	if err != nil {
// 		return err
// 	}

// 	blist := inst.raw
// 	var b1 *pem.Block = nil

// 	for _, item := range blist {
// 		if wantType.EqualString(item.Type) {
// 			b1 = item
// 			break
// 		}
// 	}

// 	if b1 == nil {
// 		return fmt.Errorf("no pem-block with type: '%s'", wantType.String())
// 	}

// 	inst.Block = *b1
// 	inst.raw = blist

// 	inst.Method = inst.GetHeader("method")
// 	strURL := inst.GetHeader("url")
// 	strParams := inst.GetHeader("params")

// 	u1, err := url.Parse(strURL)
// 	if err == nil {
// 		inst.PathWant = u1.Path
// 		inst.Query = inst.innerGetQueryMap(u1)
// 	}

// 	u2, err := url.Parse(strParams)
// 	if err == nil {
// 		inst.Params = inst.innerGetQueryMap(u2)
// 	}

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type Response struct {

	// Message.Type = "PEMAPI RESPONSE"
	Message

	Status Status
	Error  error
}

// [弃用]
// func (inst *Response) Text() ResponseText {

// 	b1 := inst.innerMakeResponseBlock()
// 	inst.AddBlock(b1)
// 	inst.RemoveRepeatedBlocks()

// 	str, err := inst.EncodeString()
// 	if err != nil {
// 		str = err.Error()
// 	}

// 	return ResponseText(str)
// }

// [弃用]
// func (inst *Response) innerMakeResponseBlock() *pem.Block {

// 	now := time.Now()
// 	tstamp := now.UnixMilli()
// 	status := inst.Status
// 	body := inst.Bytes
// 	err := inst.Error

// 	if err != nil {
// 		status.Code = http.StatusInternalServerError
// 		inst.SetHeader("error", err.Error())
// 	}
// 	if status.Code == 0 {
// 		status.Code = http.StatusOK
// 	}

// 	inst.SetHeader("status", status.String())
// 	inst.SetHeader("date", now.String())
// 	inst.SetHeader("timestamp", strconv.FormatInt(tstamp, 10))
// 	inst.SetHeader("id", "response")
// 	inst.SetContentLength(len(body))

// 	block := new(pem.Block)
// 	*block = inst.Block
// 	block.Type = BlockTypeResponse.String()
// 	return block
// }

// [弃用]
// func (inst *Response) Parse(txt ResponseText) error {

// 	const wantType = BlockTypeResponse

// 	str := txt.String()
// 	err := inst.Message.DecodeString(str)
// 	if err != nil {
// 		return err
// 	}

// 	// find response block
// 	blist := inst.raw
// 	var b1 *pem.Block = nil
// 	for _, item := range blist {
// 		if wantType.EqualString(item.Type) {
// 			b1 = item
// 			break
// 		}
// 	}
// 	if b1 == nil {
// 		return fmt.Errorf("no pem-block with type: '%s'", wantType.String())
// 	}

// 	inst.Block = *b1
// 	inst.raw = blist

// 	strErr := inst.GetHeader("error")
// 	strSta := inst.GetHeader("status")
// 	status, _ := ParseStatus(strSta)

// 	if strErr != "" {
// 		inst.Error = fmt.Errorf("%s", strErr)
// 	}
// 	inst.Status = *status

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////
// EOF
