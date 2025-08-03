package test

import (
	"github.com/bwmarrin/snowflake"
	"testing"
)

func BenchmarkTid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node, err := snowflake.NewNode(1)
		if err != nil {
			panic(err)
		}
		str := node.Generate().String()
		println(str)
	}
}
