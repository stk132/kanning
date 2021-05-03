package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type metrics struct {
	name string
	value float64
}

func newMetrics(name string, value float64) *metrics {
	return &metrics{
		name:  name,
		value: value,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	metricsList := []*metrics{}
	for scanner.Scan() {
		metric, err := parse(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		metricsList = append(metricsList, metric)
	}

	m := map[string]interface{}{}

	for _, v := range metricsList {
		names := strings.Split(v.name, ".")[1:]
		size := len(names)
		currentMap := m
		for i, name := range names {
			if i == size - 1 {
				currentMap[name] = v.value
			} else {
				if _, ok := currentMap[name]; ok {
					currentMap = currentMap[name].(map[string]interface{})
				} else {
					tmp := map[string]interface{}{}
					currentMap[name] = tmp
					currentMap = tmp
				}
			}

		}
	}

	buf, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

func parse(line string) (*metrics, error) {
	items := strings.Fields(line)
	if len(items) != 3 {
		return nil, fmt.Errorf("line should contains 3 items, but %d", len(items))
	}
	
	value, err := strconv.ParseFloat(items[1], 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert float64 from %s", items[1])
	}

	return newMetrics(items[0], value), nil
}
