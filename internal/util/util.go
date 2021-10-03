package util

import (
	"encoding/json"
	"fmt"
)

func DebugPrint(prefix string, vs ...interface{}) {
	output := "------------------DEBUG-PRINT--------------------\n"
	output += fmt.Sprint(prefix, ":\n")
	for _, v := range vs {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			output += fmt.Sprintf("%#+v\n", v)
			continue
		}
		output += fmt.Sprintf("%s\n", b)
	}
	fmt.Print(output)
}
