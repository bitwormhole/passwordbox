package loggers

import "testing"

func TestLogger(t *testing.T) {

	table := make(map[string]Level)
	keys := make([]string, 0)

	table["demo-1"] = INFO
	table["demo-2"] = WARN
	table["demo-3"] = ERROR
	table["demo-4"] = FATAL
	table["demo-5"] = DEBUG
	table["demo-6"] = TRACE

	for key := range table {
		keys = append(keys, key)
	}

	SetLevelEnabled(TRACE)

	for index, key := range keys {
		level := table[key]
		logger1 := GetLogger(level)
		value := "this is a demo for level: " + level.String()
		logger1.Printf("[demo index:%d name:%s msg:%s]", index, key, value)
	}

}
