package main

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	r := rule{
		Type:    "pattern",
		Pattern: []int{4, 4, 4, 1},
	}

	data := []int{4, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1}
	res, err := r.check(data)
	fmt.Println(res)
	fmt.Println(err)

}

func TestCompostion(t *testing.T) {
	r1 := rule{
		Type:   "delta",
		Check:  ">",
		Change: 5,
		Over:   7,
	}
	r2 := rule{
		Type:   "delta",
		Check:  "=",
		Change: 5,
		Over:   7,
	}
	r3 := rule{
		Type: "composition",
		All:  []rule{r1, r2},
	}

	data := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

	fmt.Println(r1.check(data))
	fmt.Println(r2.check(data))
	fmt.Println(r3.check(data))
}
