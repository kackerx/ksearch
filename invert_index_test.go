package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestBuildInvertIndex(t *testing.T) {
	tests := []struct {
		docs []*Doc

		wantRes map[string][]int
	}{
		{
			docs: []*Doc{
				{1, []string{"go", "数据结构"}},
				{2, []string{"go", "数据库"}},
			},
			wantRes: map[string][]int{
				"go":   {1, 2},
				"数据结构": {1},
				"数据库":  {2},
			},
		},
	}

	for _, tt := range tests {
		t.Run("df", func(t *testing.T) {
			idxMap := BuildInvertIndex(tt.docs)
			if !reflect.DeepEqual(idxMap, tt.wantRes) {
				t.Error("获取结果错误", idxMap)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	itor := Split("i am man")
	for i, s := range itor {
		fmt.Println(i, s)
		break
	}

	for i, s := range itor {
		fmt.Println(i, s)
	}
}

func Split(s string) func(yield func(k int, v string) bool) {
	parts := strings.Split(s, " ")
	return func(yield func(k int, v string) bool) {
		for i, part := range parts {
			if !yield(i, part) {
				return
			}
		}
	}
}

func TestIterator(t *testing.T) {
	for i, v := range For {
		fmt.Println(i, v)
		if i == 3 {
			break
		}
	}

	for i := range 3 {
		fmt.Println(i)
	}
}

func For(yield func(k, v int) bool) {
	for i := 0; i < 5; i++ {
		if !yield(i, i+2) {
			fmt.Println("中断", i)
			return
		}
		if i == 4 {
			break
		}
	}
}
