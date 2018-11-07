package mergesort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const size = 100000

var ss [][]int

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	ss = make([][]int, 0, 1000)
	for i := 0; i < cap(ss); i++ {
		s := make([]int, size)
		for i := 0; i < size; i++ {
			s[i] = rand.Intn(size)
		}
		ss = append(ss, s)
	}
}

func TestMergesort(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	mergesort(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func TestParallelMergesort1(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	parallelMergesort1(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func TestParallelMergesort2(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	parallelMergesort2(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func TestParallelMergesort3(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	parallelMergesort3(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func BenchmarkMergesort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergesort(ss[i%len(ss)])
	}
}

func BenchmarkParallelMergesort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelMergesort1(ss[i%len(ss)])
	}
}

func BenchmarkParallelMergesort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelMergesort2(ss[i%len(ss)])
	}
}

func BenchmarkParallelMergesort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelMergesort3(ss[i%len(ss)])
	}
}
