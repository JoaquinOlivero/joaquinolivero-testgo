package main

import (
	"reflect"
	"testing"
)

func TestZeroSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  []int
	}{
		{"Should be 5, 8", []int{3, 4, -7, 5, -6, 2, 5, -1, 8}, []int{5, 8}},
		{"Should be 1, 2, 1", []int{1, 2, -3, 3, 1}, []int{1, 2, 1}},
		{"Should be 2, 3", []int{1, -1, 2, 3}, []int{2, 3}},
		{"Should be 1,2,3,4", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"Should be 6", []int{1, 2, -3, 4, -4, 5, -5, 6}, []int{6}},
		{"Should be 7", []int{7}, []int{7}},
		{"Should be empty", []int{0, 0, 1, -1, 2, -2}, []int{}},
		{"Should be empty", []int{1, 2, -3}, []int{}},
		{"Should be empty", []int{1, 2, 3, -3, -2, -1}, []int{}},
		{"Should be empty", []int{}, []int{}},
		{"Should be empty", nil, []int{}},
	}

	// loop through test tables
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ZeroSum(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestEncryptMessage(t *testing.T) {
	var tests = []struct {
		key  string
		msg  string
		want string
	}{
		{"dcj", "I love prOgrAmming", "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing"},
		{"abc", "Hello World!", "Habcellabco Wabcorld!"},
		{"xyz", "Golang is Awesome!", "Gxyzolxyzang xyzis xyzAwxyzesxyzomxyze!"},
		{"", "I love coding", "dcjI ldcjovdcje cdcjoddcjing"},
		{"abc", "", ""},
		{"", "", ""},
		{"", "Hello World!", "Hdcjelldcjo Wdcjorld!"},
		{"key", "bcdfg", "bcdfg"},
		{"xyz", "aeiouAEIOU", "xyzaxyzexyzixyzoxyzuxyzAxyzExyzIxyzOxyzU"},
		{"sec", "H@ppy c0ding!", "H@ppy c0dsecing!"},
	}

	for _, tt := range tests {
		result := EncryptMessage(Params{tt.key, tt.msg})
		if result != tt.want {
			t.Errorf("EncryptMessage(%q, %q) = %q; want %q", tt.key, tt.msg, result, tt.want)
		}
	}
}
