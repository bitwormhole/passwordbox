package properties

import "sort"

////////////////////////////////////////////////////////////////////////////////

type Table interface {
	Put(name, value string)

	Get(name string) string

	Keys() []string

	Reset()

	Trim()

	Count() int

	Import(map[string]string)

	Export(out map[string]string) map[string]string
}

////////////////////////////////////////////////////////////////////////////////

type innerTable struct {
	tab map[string]string
}

func NewTable() Table {
	return new(innerTable)
}

func (inst *innerTable) Put(name, value string) {
	tab := inst.innerGetRawTable(true)
	tab[name] = value
}

func (inst *innerTable) Get(name string) string {
	tab := inst.innerGetRawTable(false)
	if tab == nil {
		return ""
	}
	return tab[name]
}

func (inst *innerTable) Keys() []string {
	dst := make([]string, 0)
	src := inst.innerGetRawTable(false)
	for key := range src {
		dst = append(dst, key)
	}
	sort.Strings(dst)
	return dst
}

func (inst *innerTable) Reset() {
	inst.tab = make(map[string]string)
}

func (inst *innerTable) Trim() {
	dst := make(map[string]string)
	src := inst.innerGetRawTable(false)
	for key, value := range src {
		if value == "" {
			continue
		}
		dst[key] = value
	}
	inst.tab = dst
}

func (inst *innerTable) Count() int {
	count := len(inst.tab)
	return count
}

func (inst *innerTable) Import(src map[string]string) {
	dst := inst.innerGetRawTable(true)
	for key, value := range src {
		if value == "" {
			continue
		}
		dst[key] = value
	}
	inst.tab = dst
}

func (inst *innerTable) Export(dst map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	src := inst.innerGetRawTable(false)
	for key, value := range src {
		if value == "" {
			continue
		}
		dst[key] = value
	}
	return dst
}

func (inst *innerTable) innerGetRawTable(create bool) map[string]string {
	tab := inst.tab
	if create {
		if tab == nil {
			tab = make(map[string]string)
			inst.tab = tab
		}
	}
	return tab
}

////////////////////////////////////////////////////////////////////////////////
