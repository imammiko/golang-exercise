package main

import (
	"fmt"
)

func main() {
	fmt.Println(

		checkSallary([4][6]interface{}{{"H", "H", "H", "H", "H", "H"}, {[]string{"A", "B"}, []interface{}{"l", 4}, "H", []interface{}{"L", 3}, "H", "H"}, {"H", "H", "H", "T", []string{"A", "C"}, []string{"A", "C"}}, {"H", "T", []string{"A", "S"}, "H", "H", "H"}}),
	)

}

func checkSallary(date [4][6]interface{}) int {
	// fmt.Println(string(reflect.TypeOf(date[1][0])) == "[]
	// string")
	var gapok int
	gapok = 5000000
	for _, v := range date {
		for _, w := range v {

			switch w.(type) {
			case []string:
				if w.([]string)[0] == "A" {
					switch w.([]string)[1] {
					case "S":
						continue
					case "C":
						continue
					case "B":
						gapok = gapok - 200000
					}
				}

			case []interface{}:
				if w.([]interface{})[0] == "L" {
					gapok = gapok - (50000 * (w.([]interface{})[1].(int)))
				}

			case string:
				if w == "H" {
					continue
				} else if w == "T" {
					gapok = gapok - 50000
				}
			}
			fmt.Println(w)
		}
	}

	return gapok

}
