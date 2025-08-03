package test

import (
	"fmt"
	"github.com/btnguyen2k/olaf"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/jakehl/goid"
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

func BenchmarkTid2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v4UUID := goid.NewV4UUID()
		fmt.Println(v4UUID)
	}
}

func BenchmarkTid3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := olaf.NewOlaf(1981)
		id64 := o.Id64()
		id64Hex := o.Id64Hex()
		id64Ascii := o.Id64Ascii()
		fmt.Println("ID 64-bit (int)   : ", id64, " / Timestamp: ", o.ExtractTime64(id64))
		fmt.Println("ID 64-bit (hex)   : ", id64Hex, " / Timestamp: ", o.ExtractTime64Hex(id64Hex))
		fmt.Println("ID 64-bit (ascii) : ", id64Ascii, " / Timestamp: ", o.ExtractTime64Ascii(id64Ascii))

	}
}

func BenchmarkUuid1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := uuid.NewString()
		println(str)
	}
}
