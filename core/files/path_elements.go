package files

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type PathElement string

func (pe PathElement) String() string {
	return string(pe)
}

func (pe PathElement) Normalize() PathElement {
	str := pe.String()
	str = strings.TrimSpace(str)
	return PathElement(str)
}

////////////////////////////////////////////////////////////////////////////////

type PathElementArray []PathElement

func (array PathElementArray) Normalize() (PathElementArray, error) {
	tmp := make(PathElementArray, 0)
	for _, item := range array {
		item2 := item.Normalize()
		if item2 == "" {
			continue // ignore
		} else if item2 == "." {
			continue // ignore
		} else if item2 == ".." {
			parent, err := tmp.innerGetParent()
			if err != nil {
				return nil, err
			}
			tmp = parent
		} else if item2 == "~" {
			return nil, fmt.Errorf("unsupported path element : '~'")
		} else {
			tmp = append(tmp, item2)
		}
	}
	return tmp, nil
}

func (array PathElementArray) innerGetParent() (PathElementArray, error) {
	count := len(array)
	if count > 0 {
		return array[0 : count-1], nil
	}
	return nil, fmt.Errorf("PathElementArray: no parent")
}

func (array PathElementArray) IsAbsolute() bool {
	return !array.IsRelative()
}

func (array PathElementArray) IsRelative() bool {
	name1st := array.GetFirstName()
	if name1st == "." {
		return true
	} else if name1st == ".." {
		return true
	} else if name1st == "~" {
		return true
	} else if name1st == "" {
		return false
	}
	return false
}

func (array PathElementArray) String() string {
	return array.Path().String()
}

func (array PathElementArray) Path() Path {
	const sep = '/'
	abs := array.IsAbsolute()
	builder := new(strings.Builder)
	for _, item1 := range array {
		item2 := strings.TrimSpace(item1.String())
		builder.WriteRune(sep)
		builder.WriteString(item2)
	}
	str := builder.String()
	if !abs {
		str = "." + str
	}
	return Path(str)
}

// 获取元素列表中最后一个有效名称
func (array PathElementArray) GetLastName() string {
	count := len(array)
	for i := count - 1; i >= 0; i-- {
		el := array[i]
		name := strings.TrimSpace(el.String())
		if name == "" {
			continue
		}
		return name
	}
	return ""
}

// 获取元素列表中第一个有效名称
func (array PathElementArray) GetFirstName() string {
	for _, el := range array {
		name := strings.TrimSpace(el.String())
		if name == "" {
			continue
		}
		return name
	}
	return ""
}

////////////////////////////////////////////////////////////////////////////////
