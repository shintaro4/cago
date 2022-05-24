package main

import (
	"testing"
)


func TestInitRule(t *testing.T) {

	invalidRuleNumber := 256
	_, err := InitRule(uint(invalidRuleNumber))
	if err == nil {
		t.Errorf("given rule number is out of range; %d", invalidRuleNumber)
	}

	cases := []struct{
		description string
		ruleNumber uint
		expected []bool
	} {
		{
			description: "rule 0",
			ruleNumber: 0,
			expected: []bool{false, false, false, false, false, false, false, false},
		},
		{
			description: "rule 1",
			ruleNumber: 1,
			expected: []bool{true, false, false, false, false, false, false, false},
		},
		{
			description: "rule 30",
			ruleNumber: 30,
			expected: []bool{false, true, true, true, true, false, false, false},
		},
		{
			description: "rule 255",
			ruleNumber: 255,
			expected: []bool{true, true, true, true, true, true, true, true},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := InitRule(tt.ruleNumber)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d, but got %d", len(tt.expected), len(result))
			}
			for i := 0; i < len(result); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, but got %v", tt.expected[i], result[i])
				}
			}
		})
	}
}


func TestNext(t *testing.T) {

	cases := []struct{
		description string
		ruleNumber uint
		initialCells []bool
		expected []bool
	} {
		{
			description: "a cell with rule 30",
			ruleNumber: 30,
			initialCells: []bool{true},
			expected: []bool{true, true, true},
		},
	}

    for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			rule, _ := InitRule(tt.ruleNumber)
			result, _ := Next(rule, tt.initialCells)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d, but got %d", len(tt.expected), len(result))
			}
			for i := 0; i < len(result); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, but got %v", tt.expected[i], result[i])
				}
			}
		})
	}
}
