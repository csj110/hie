package main

import (
	"reflect"
	"testing"
)

var circularTests = []struct {
	in  []int
	out bool
}{
	{[]int{20, 39, 87, 0, 1}, true},
	{[]int{2, 3}, true},
	{[]int{6, 2, 3, 4, 5}, true},
	{[]int{5, 6, 54, 435, 888, 9999, -8, 1, 2}, true},
	{[]int{3, 2, 4, 5, 6}, true},
}

var inArrayTests = []struct {
	in1 []string
	in2 []string
	out []string
}{
	{
		[]string{"live", "arp", "strong"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"}, 
		[]string{"arp", "live", "strong"},
	},
	{
		[]string{"cod", "code", "wars", "ewar", "ar"},
		[]string{},
		[]string{},
	},
}

var bitCountingTests=[]struct{
	in uint
	out int
}{
	{0,0},
	{4,1},
	{7,3},
	{9,2},
	{10,2},
}

func TestCircular(t *testing.T) {
	for _, test := range circularTests {
		result := isCircularSorted(test.in)
		if result != test.out {
			t.Errorf("Expected %t, got %t", test.out, result)
		}
	}
}

func TestInArray(t *testing.T) {
	for _, test := range inArrayTests {
		result := InArray(test.in1, test.in2)
		if !reflect.DeepEqual(test,result) {
			t.Errorf("Expected %v,got %v", test.out, result)
		}
	}
}


func TestBitCounting(t *testing.T){
	for _, test := range bitCountingTests {
		if res:=CountBits(test.in);res!=test.out{
			t.Errorf("in:%v,get:%v,expected:%v",test.in,res,test.out)
		}
	}
}
