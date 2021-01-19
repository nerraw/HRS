package main

import "errors"

type rule struct {
	ID      int    `json:"id"`
	Name    string `json:"rule-name"`
	Type    string `json:"type"`
	Value   int    `json:"value"`
	Check   string `json:"check"`
	Change  int    `json:"change"`
	Over    int    `json:"over"`
	All     []rule `json:"all"`
	Pattern []int  `json:"pattern"`
}

func gt(a, b int) bool {
	return a > b
}

func lt(a, b int) bool {
	return a < b
}

func eq(a, b int) bool {
	return a == b
}

func getCompare(checkName string) (func(a, b int) bool, error) {
	switch checkName {
	case ">":
		return gt, nil
	case "<":
		return lt, nil
	case "=":
		return eq, nil
	default:
		return nil, errors.New("unknown comparison type " + checkName)
	}
}

func (r rule) check(data []int) ([]bool, error) {
	res := make([]bool, len(data))
	switch r.Type {
	case "comparison":
		compare, err := getCompare(r.Check)
		if err != nil {
			return nil, err
		}
		for i, val := range data {
			res[i] = compare(val, r.Value)
		}
		return res, nil

	case "delta":
		compare, err := getCompare(r.Check)
		if err != nil {
			return nil, err
		}
		for i := r.Over; i < len(data); i++ {
			res[i] = compare(data[i]-data[i-r.Over], r.Change)
		}
		return res, nil
	case "pattern":
		if len(r.Pattern) < 1 {
			return res, nil
		}
		for i := len(r.Pattern) - 1; i < len(data); i++ {
			p := true
			for j, val := range r.Pattern {
				p = p && ((data[i-len(r.Pattern)+j+1] & val) == val)
			}
			res[i] = p
		}
		return res, nil
	case "composition":
		for _, subRule := range r.All {
			newRes, err := subRule.check(data)
			if err != nil {
				return nil, err
			}
			for j := range newRes {
				res[j] = res[j] || newRes[j]
			}
		}
		return res, nil
	}

	return nil, errors.New("unknown rule type " + r.Type)
}
