package test

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"testing"
)

var name = []string{"1", "2", "3", "4"}

func BenchmarkGetVedioLikeList(b *testing.B) {
	for _, i := range name {
		output := service.Add(i, i+"2")
		if output != 1 {
			b.Errorf("output: %d", output)
		}
		fmt.Println(output)
	}
}
