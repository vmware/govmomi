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
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"text/template"
)

func load(name string) string {
	s, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(s)
}

func TestReadCounterInfo(t *testing.T) {
	res := load("fixtures/counterinfo.dsv")
	s := ParseCounterInfo(res)
	expect := CommandParameterType{
		Name: "LocalMemoryInKB",
		Type: "U32",
	}

	if !reflect.DeepEqual(s["SchedGroup"].Types[24], expect) {
		t.Errorf("got %v", s["SchedGroup"].Types[24])
	}
}

func TestReadCounterStats(t *testing.T) {
	res := load("fixtures/counterinfo.dsv")
	counterInfo := ParseCounterInfo(res)
	statsFile := load("fixtures/fetchstats.dsv")
	stats := ParseCounterValues(statsFile, counterInfo)
	expect := CounterValue{
		TypeName: "SchedGroup",
		Values: map[string]interface{}{
			"MemAllocMax":            -1,
			"MemAllocShares":         -3,
			"IsValid":                true,
			"CPUAllocMin":            0,
			"CPUAllocUnitsStr":       "mhz",
			"LocalMemoryInKB":        29102080,
			"CPUAllocMax":            -1,
			"CPUAllocUnits":          1,
			"HomeNodes":              3,
			"NumOfCPUClients":        9,
			"LocalityPct":            100,
			"CPUAllocShares":         -3,
			"MemAllocUnits":          4,
			"GroupName":              "vm.13713826",
			"VMName":                 "frigo001",
			"CPUAllocMinLimit":       -1,
			"NumOfLocalitySwap":      0,
			"GroupID":                90356474,
			"IsVM":                   true,
			"TimerPeriod":            2656,
			"MemAllocMin":            0,
			"MemAllocMinLimit":       -1,
			"MemAllocUnitsStr":       "kb",
			"IsNUMAValid":            true,
			"PowerInWatt":            6,
			"NumOfNUMANodes":         2,
			"HasMemClient":           true,
			"NumOfBalanceMigrations": 0,
			"NumOfLoadSwap":          0,
			"RemoteMemoryInKB":       0,
			"NumOfHiddenWorlds":      42,
			"GroupLeaderID":          13713826,
		},
	}

	if !reflect.DeepEqual(stats[3060], expect) {
		t.Errorf("got %v", stats[3060])
	}
}

func TestTemplatedOutput(t *testing.T) {
	res := load("fixtures/counterinfo.dsv")
	counterInfo := ParseCounterInfo(res)
	statsFile := load("fixtures/fetchstats.dsv")
	stats := ParseCounterValues(statsFile, counterInfo)
	tmpl, err := template.New("test1").Parse(DefaultTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, stats)
	if err != nil {
		panic(err)
	}
}
