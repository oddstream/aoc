package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed "input.json"
var input []byte

var dumpTotal float64

func dumpJSON(v interface{}, kn string) {
	iterMap := func(x map[string]interface{}, root string) {
		var knf string
		if root == "root" {
			knf = "%q:%q"
		} else {
			knf = "%s:%q"
		}
		var containsRed bool
		for _, v := range x {
			switch vv := v.(type) {
			case string:
				if vv == "red" {
					containsRed = true
					// fmt.Println(vv, "MAP")
				}
			}
		}
		if !containsRed {
			for k, v := range x {
				dumpJSON(v, fmt.Sprintf(knf, root, k))
			}
		}
	}

	iterSlice := func(x []interface{}, root string) {
		var knf string
		if root == "root" {
			knf = "%q:[%d]"
		} else {
			knf = "%s:[%d]"
		}
		for k, v := range x {
			dumpJSON(v, fmt.Sprintf(knf, root, k))
		}
	}

	switch vv := v.(type) {
	case string:
		// fmt.Printf("%s => (string) %q\n", kn, vv)
	case bool:
		// fmt.Printf("%s => (bool) %v\n", kn, vv)
	case float64:
		// fmt.Printf("%s => (float64) %f\n", kn, vv)
		dumpTotal += vv
	case map[string]interface{}:
		// fmt.Printf("%s => (map[string]interface{}) ...\n", kn)
		iterMap(vv, kn)
	case []interface{}:
		// fmt.Printf("%s => ([]interface{}) ...\n", kn)
		iterSlice(vv, kn)
	default:
		fmt.Printf("%s => (unknown?) ...\n", kn)
	}
}

func main() {
	// do part 1 by scanning bytes, looking for numbers not in quotes
	var total int = 0
	re := regexp.MustCompile(`[0-9]+`)
	found := re.FindAllIndex(input, -1)
	for _, pair := range found {
		p0 := pair[0]
		p1 := pair[1]
		// there are no -ve numbers in quotes
		// regexp matches greedily
		if input[p0-1] != byte('"') {
			str := string(input[p0:p1])
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(str, err)
			}
			if input[p0-1] == '-' {
				num = -num
			}
			// fmt.Println(pos, num)
			total += num
		}
	}
	fmt.Println("part 1", total) // 156366

	// do part 2 by loading and parsing json
	// https://stackoverflow.com/questions/30341588/how-to-parse-a-complicated-json-with-go-unmarshal
	var obj []interface{}
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		fmt.Println(err)
	} else {
		dumpJSON(obj, "root")
	}
	fmt.Println("part 2", dumpTotal) // 96852
}
