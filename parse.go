package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func parseData(fname string) ([]int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	res := []int{}

	// Parse optional header
	record, err := r.Read()
	if err != nil {
		return nil, err
	}
	i, err := strconv.Atoi(record[0])
	if err == nil {
		res = append(res, i)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) != 1 {
			return nil, fmt.Errorf("invalid number of data points: %v", record)
		}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, nil
}

func parseRules(fname string) ([]rule, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	js, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	for len(js) > 0 && js[0] != '{' && js[0] != '[' {
		js = js[1:]
	}

	rules := []rule{}
	err = json.Unmarshal([]byte(js), &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
