package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const RuleSequenceSize = 8

// returns a boolean list expression of the ruleNumber.
func InitRule(ruleNumber uint) ([]bool, error) {
	if ruleNumber > 255 {
		return nil, errors.New("invalid rule number")
	}

	var r = make([]bool, RuleSequenceSize)
	for i := 0; i < RuleSequenceSize; i++ {
		if ruleNumber&(1<<i) > 0 {
			r[i] = true
		} else {
			r[i] = false
		}
	}

	return r, nil
}

// returns the next generation of the cells
func Next(rule []bool, cells []bool) ([]bool, error) {
	f := func(i int) bool {
		var l, c, r uint = 0, 0, 0
		if 0 <= i-1 && i-1 < len(cells) && cells[(i-1)%len(cells)] {
			l = 4
		}
		if 0 <= i && i < len(cells) && cells[i] {
			c = 2
		}
		if 0 <= i+1 && i+1 < len(cells) && cells[(i+1)%len(cells)] {
			r = 1
		}
		return rule[l+c+r]
	}

	var r = make([]bool, len(cells)+2)
	for i := -1; i < len(cells)+1; i++ {
		r[i+1] = f(i)
	}

	return r, nil
}

// print cells nicely
func PrintCells(cells []bool, n int) {
	s := make([]string, len(cells))
	for i, c := range cells {
		if c {
			s[i] = "o"
		} else {
			s[i] = " "
		}
	}
	padding := strings.Repeat(" ", n)
	fmt.Println(padding + strings.Join(s, ""))
}

func main() {
	ruleNumber := flag.Uint("rule", 110, "a rule number")
	n := flag.Int("n", 16, "the number of the generations")
	flag.Parse()

	if *n <= 0 {
		fmt.Println("'n' must be a positive integer.")
		os.Exit(0)
	}

	cells := []bool{true}
	rule, err := InitRule(uint(*ruleNumber))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for i := 0; i < *n; i++ {
		PrintCells(cells, *n-i)
		cells, _ = Next(rule, cells)
	}

}
