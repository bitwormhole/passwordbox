package properties

import "sort"

type Table struct {
	tab map[string]string
}

func NewTable() *Table {
	return new(Table)
}

func (inst *Table) Put(name, value string) {
	tab := inst.innerGetRawTable(true)
	tab[name] = value
}

func (inst *Table) Get(name string) string {
	tab := inst.innerGetRawTable(false)
	if tab == nil {
		return ""
	}
	return tab[name]
}

func (inst *Table) Keys() []string {
	dst := make([]string, 0)
	src := inst.innerGetRawTable(false)
	for key := range src {
		dst = append(dst, key)
	}
	sort.Strings(dst)
	return dst
}

func (inst *Table) Reset() {
	inst.tab = make(map[string]string)
}

func (inst *Table) Trim() {
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

func (inst *Table) Count() int {
	count := len(inst.tab)
	return count
}

func (inst *Table) Import(src map[string]string) {
	dst := inst.innerGetRawTable(true)
	for key, value := range src {
		if value == "" {
			continue
		}
		dst[key] = value
	}
	inst.tab = dst
}

func (inst *Table) Export(dst map[string]string) map[string]string {
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

func (inst *Table) innerGetRawTable(create bool) map[string]string {
	tab := inst.tab
	if create {
		if tab == nil {
			tab = make(map[string]string)
			inst.tab = tab
		}
	}
	return tab
}
