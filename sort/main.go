package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int)
	m["CDM"] = 1
	m["backup"] = 2
	m["dg"] = 3
	m["resource"] = 4
	m["storage"] = 5
	m["system"] = 6
	fmt.Println(m)
	temp := []string{"backup", "CDM", "storage", "resource", "system", "dg"}
	sort.Slice(temp, func(i, j int) bool {
		return m[temp[i]] < m[temp[j]]
	})
	fmt.Println(temp)
}
