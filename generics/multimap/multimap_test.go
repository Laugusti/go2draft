package multimap

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	multiMap := testMultiMap()
	tests := []struct {
		key int
		want []string
	} {
		{0, []string{"zero"}},
		{1, []string{"one"}},
		{5, []string(nil)},
	}

	for _, test := range tests {
		got := multiMap.Get(test.key)
		check(t, "TestGet", test.key, test.want, got)
	}
}

func TestSet(t *testing.T) {
	multiMap := testMultiMap()
	tests := []struct {
		key int
		value string
		want []string
	} {
		{1,"uno", []string{"uno"}},
		{3,"3", []string{"3"}},
		{5,"five", []string{"five"}},
	}

	for _, test := range tests {
		multiMap.Set(test.key, test.value)
		check(t, "TestSet", test.key, test.want, multiMap.m[test.key])
	}
}

func TestAdd(t *testing.T) {
	multiMap := testMultiMap()
	tests := []struct {
		key int
		value string
		want []string
	} {
		{1,"uno", []string{"one", "uno"}},
		{3,"3", []string{"three", "3"}},
		{5,"five", []string{"five"}},
	}

	for _, test := range tests {
		multiMap.Add(test.key, test.value)
		check(t, "TestAdd", test.key, test.want, multiMap.m[test.key])
	}
}

func TestAddAll(t *testing.T) {
	multiMap := testMultiMap()
	tests := []struct {
		key int
		value []string
		want []string
	} {
		{1,[]string{"uno"}, []string{"one", "uno"}},
		{3,[]string{"tres", "3"}, []string{"three", "tres", "3"}},
		{5,[]string{"five"}, []string{"five"}},
	}

	for _, test := range tests {
		multiMap.AddAll(test.key, test.value)
		check(t, "TestAddAll", test.key, test.want, multiMap.m[test.key])
	}
}

func TestNewMap(t *testing.T) {
	multiMap := NewMap(string, string)()
	if multiMap == nil {
		t.Errorf("TestNewMap => failed to create map")
	}
	// inner map must be assignable
	multiMap.m[""] = ""
}

func TestMapOf(t *testing.T) {
	multiMap := MapOf(int, string)(func(i int, v string) { return i }, "zero", "one", "two", "three", "four")
	want := map[int][]string{
		0: []string{"zero"},
		1: []string{"one"},
		2: []string{"two"},
		3: []string{"three"},
		4: []string{"four"},
	}

	check(t, "TestMapOf", multiMap, want, multiMap.m)
}

func TestAsRawMap(t *testing.T) {
	multiMap := testMultiMap()
	want := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four"}

	rawMap := multiMap.AsRawMap(func(k int, v []string) { return v[0] })

	check(t, "TestAsRawMap", multiMap, want, rawMap)
}

func testMultiMap(type K, V)() *MultiMap(K, V) {
	return MultiMap{map[int][]string{
		0: []string{"zero"},
		1: []string{"one"},
		2: []string{"two"},
		3: []string{"three"},
		4: []string{"four"},
	}}
}

func check(t *testing.T, name string, input, expected, actual interface{}) {
	if !reflect.DeepEqual(want, rawMap) {
		t.Errorf("%s => input: %v, expected: %v, got: %v", name, input, expected, actual)
	}
}
