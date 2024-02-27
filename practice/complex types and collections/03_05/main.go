package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Map")
	states := make(map[string]string) // unorder collection
	fmt.Println(states)

	states["WA"] = "Washiton"
	states["NY"] = "NewYork"

	states["OR"] = "Oregan"
	states["CA"] = "Carifonia"
	fmt.Println(states)

	newyork := states["NY"]
	fmt.Println(newyork)

	delete(states, "WA")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v \n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Printf(states[keys[i]])
	}

}
