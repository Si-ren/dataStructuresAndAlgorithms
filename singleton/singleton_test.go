package singleton_test

import (
	"dataStructuresAndAlgorithms/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleton(t *testing.T) {
	assert := assert.New(t)
	s01 := singleton.GetSingleton()
	s02 := singleton.GetSingleton()
	assert.Equal(s01, s02, "these should be same.")
}

// testing.B是压力测试,函数名要以Benchmark开头
func BenchmarkGetSingletonBench(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetSingleton() != singleton.GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
