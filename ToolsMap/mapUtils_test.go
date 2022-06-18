package ToolsMap

import "testing"

func TestKeysAndValues(t *testing.T) {
	test := map[int]string{1: "aa", 2: "bb", 3: "cc"}
	keys, values := KeysAndValues(test)
	t.Logf("%v  %v", keys, values)
}
