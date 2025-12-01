package properties

import (
	"fmt"
	"strings"
)

func Encode(t Table) []byte {
	encoder := new(innerEncoder)
	return encoder.encodeBin(t)
}

func Format(t Table) string {
	encoder := new(innerEncoder)
	return encoder.encodeStr(t)
}

func Decode(raw []byte, t Table) error {
	decoder := new(innerDecoder)
	return decoder.decodeBin(raw, t)
}

func Parse(raw string, t Table) error {
	decoder := new(innerDecoder)
	return decoder.decodeStr(raw, t)
}

////////////////////////////////////////////////////////////////////////////////

type innerEncoder struct {
}

func (inst *innerEncoder) encodeBin(t Table) []byte {
	str := inst.encodeStr(t)
	return []byte(str)
}

func (inst *innerEncoder) encodeStr(t Table) string {
	builder := new(strings.Builder)
	keys := t.Keys()
	for _, key := range keys {
		val := t.Get(key)
		builder.WriteString(key)
		builder.WriteString(" = ")
		builder.WriteString(val)
		builder.WriteByte('\n')
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type innerDecoder struct {
}

func (inst *innerDecoder) decodeBin(b []byte, out Table) error {
	str := string(b)
	return inst.decodeStr(str, out)
}

func (inst *innerDecoder) decodeStr(str string, out Table) error {
	if out == nil {
		return fmt.Errorf("param: out is nil")
	}
	rows := inst.splitToRows(str)
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue // ignore
		} else if strings.HasPrefix(row, "#") {
			continue // ignore
		}
		k, v, err := inst.parseKeyValue(row)
		if err != nil {
			return err
		}
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		out.Put(k, v)
	}
	return nil
}

func (inst *innerDecoder) splitToRows(str string) []string {
	const (
		sep1 = "\r"
		sep2 = "\n"
		sep  = sep2
	)
	str = strings.ReplaceAll(str, sep1, sep2)
	return strings.Split(str, sep)
}

func (inst *innerDecoder) parseKeyValue(str string) (key string, value string, err error) {
	const (
		sep = "="
		n   = 2
	)
	array := strings.SplitN(str, sep, n)
	if len(array) == n {
		key = array[0]
		value = array[1]
		err = nil
		return
	}
	return "", "", fmt.Errorf("bad format of 'key-value' row: %s", str)
}

////////////////////////////////////////////////////////////////////////////////
