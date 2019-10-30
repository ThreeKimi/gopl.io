package popcount_test

import "testing"
import "gopl.io/ch2/popcount"

func BenchmarkPopCount1(b *testing.B){
  for i := 0; i < b.N; i++ {
    popcount.PopCount(0xff2345ff)
  }
}

func BenchmarkPopCount23(b *testing.B){
  for i := 0; i < b.N; i++ {
    popcount.PopCount23(0xff2345ff)
  }
}

func BenchmarkPopCount24(b *testing.B){
  for i := 0; i < b.N; i++ {
    popcount.PopCount24(0xff2345ff)
  }
}

func BenchmarkPopCount25(b *testing.B){
  for i := 0; i < b.N; i++ {
    popcount.PopCount25(0xff2345ff)
  }
}

