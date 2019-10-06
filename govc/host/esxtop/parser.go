/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package esxtop

import (
	"strconv"
	"strings"
)

type CommandParameterType struct {
	Name string
	Type string
}

type CommandParameters struct {
	Types []CommandParameterType
}

func ParseCounterInfo(s string) (r map[string]CommandParameters) {
	r = make(map[string]CommandParameters)
	for _, line := range strings.Split(s, "\n") {
		var commandParameters CommandParameters
		var name string
		for i, val := range strings.Split(strings.TrimRight(line, "\r\n"), "|") {
			switch {
			case i == 1:
				name = val
			case i > 1:
				pt := strings.Split(val, ",")
				if len(pt) > 1 {
					commandParameters.Types = append(commandParameters.Types, CommandParameterType{pt[0], pt[1]})
				} else {
					// stop parsing this line there
					break
				}
			}
		}
		if name != "" {
			r[name] = commandParameters
		}
	}
	return r
}

type CounterValue struct {
	TypeName string
	Values   map[string]interface{}
}

func Convert(val string, t string) (r interface{}) {
	switch t {
	case "U32", "U64", "S32", "S64":
		r, _ = strconv.Atoi(val)
	case "B":
		r, _ = strconv.ParseBool(val)
	default:
		r = val
	}
	return r
}

func ParseCounterValueLine(l string, p map[string]CommandParameters) (v CounterValue) {
	for i, val := range strings.Split(l, "|") {
		switch {
		case i == 1:
			v.TypeName = val
			v.Values = make(map[string]interface{}, 0)
		case i > 1:
			if val != "" {
				pt := p[v.TypeName].Types[i-2]
				v.Values[pt.Name] = Convert(val, pt.Type)
			}
		}
	}
	return v
}

const (
	COUNTER_VALUE = "==COUNTER-VALUE=="
)

func ParseCounterValues(s string, p map[string]CommandParameters) (v []CounterValue) {
	var inCounterValues bool = false
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimRight(line, "\r\n")
		switch {
		case line == COUNTER_VALUE:
			inCounterValues = true
		case strings.HasPrefix(line, "=="):
			inCounterValues = false
		case inCounterValues:
			v = append(v, ParseCounterValueLine(line, p))
		}
	}
	return v
}
