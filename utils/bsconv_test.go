package utils

import "testing"

var s = "adsfasdfadsfadsfasdfadfadfasdfasdfadsfasdfasdfasdfsadfas"

//func BenchmarkB2sNew(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		_ = Scb(s)
//	}
//}
//func BenchmarkB2sNormal(b *testing.B) {
//	var _ []byte
//	for i := 0; i < b.N; i++ {
//		_ = []byte(s)
//	}
//}

var (
	bt       = []byte("adsfasdfadsfadsfasdfadfadfasdfasdfadsfasdfasdfasdfsadfas")
	cha byte = 'a'
)

const concha byte = 'a'

//func BenchmarkS2BNew(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		_ = Bcs(bt)
//	}
//}
//func BenchmarkS2BNormal(b *testing.B) {
//	var _ []byte
//	for i := 0; i < b.N; i++ {
//		_ = string(bt)
//	}
//}

func BenchmarkBcsChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(concha)
	}
}

func BenchmarkBcsCharNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BcsChar(concha)
	}
}
