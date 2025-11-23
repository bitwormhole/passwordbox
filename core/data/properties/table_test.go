package properties

import "testing"

func TestTableCount(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()
	table2 := itt.createNewTable()

	count1 := table1.Count()
	count2 := table2.Count()

	t.Logf("count(table1) = %d", count1)
	t.Logf("count(table2) = %d", count2)
}

func TestTableKeys(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()

	keys := table1.Keys()
	for idx, key := range keys {
		t.Logf("  key[%d] : %s", idx, key)
	}

}

func TestTableReset(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()

	itt.logItems("before", table1, t)

	table1.Reset()

	itt.logItems("after", table1, t)

}

func TestTableTrim(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()

	table1.Put("kiwi", "")
	table1.Put("mango", "")
	table1.Put("peach", "")

	itt.logItems("before", table1, t)

	table1.Trim()

	itt.logItems("after", table1, t)

}

func TestTableGetPut(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()
	table2 := itt.createNewTable()

	ids := []string{
		"kiwi", "mango", "peach",
	}

	for _, id := range ids {
		val := table1.Get(id)
		table2.Put(id, val)
	}

	itt.logItems("src", table1, t)
	itt.logItems("dst", table2, t)

}

func TestTableExportImport(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()
	table2 := itt.createNewTable()

	tmp := table1.Export(nil)
	table2.Import(tmp)

	itt.logItems("src", table1, t)
	itt.logItems("dst", table2, t)

}

func TestFormatTable(t *testing.T) {

	itt := new(innerTableTester)
	table := itt.createNewTableWithData()

	str := Format(table)

	t.Logf("format(table) : \n%s", str)
}

func TestParseTable(t *testing.T) {

	itt := new(innerTableTester)
	table1 := itt.createNewTableWithData()
	table2 := itt.createNewTable()

	str1 := Format(table1)

	err := Parse(str1, table2)
	if err != nil {
		t.Error(err)
	}

	str2 := Format(table2)

	t.Logf("format(table1) : \n%s", str1)
	t.Logf("format(table2) : \n%s", str2)
}

////////////////////////////////////////////////////////////////////////////////

type innerTableTester struct {
}

func (inst *innerTableTester) createNewTable() *Table {
	table := NewTable()
	return table
}

func (inst *innerTableTester) createNewTableWithData() *Table {

	src := make(map[string]string)

	src["blueberry"] = "蓝莓"
	src["kiwi"] = "猕猴桃"
	src["coconut"] = "椰子"
	src["lemon"] = "柠檬"
	src["cherry"] = "樱桃"
	src["mango"] = "芒果"
	src["pineapple"] = "菠萝"
	src["pear"] = "梨子"
	src["peach"] = "桃子"
	src["watermelon"] = "西瓜"
	src["strawberry"] = "草莓"

	table := NewTable()
	table.Import(src)
	return table
}

func (inst *innerTableTester) logItems(tag string, table *Table, t *testing.T) {
	str := Format(table)
	t.Logf("%s\n%s", tag, str)
}
