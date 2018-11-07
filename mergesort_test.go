package mergesort

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const size = 100000

var ss [][]int

func init() {
	var s []int
	ss = make([][]int, 0, 100)

	f, _ := os.Open("testcases.txt")
	bf := bufio.NewReader(f)
	for {
		line, _ := bf.ReadString('\n')
		if len(line) == 0 {
			break
		}
		json.Unmarshal([]byte(line), &s)
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
	helper := make([]int, len(s))
	parallelMergesort2(s, helper)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func TestParallelMergesort3(t *testing.T) {
	s := []int{5, 8, 9, 5, 0, 10, 1, 6}
	parallelMergesort3(s)
	assert.Equal(t, []int{0, 1, 5, 5, 6, 8, 9, 10}, s)
}

func BenchmarkMergesort(b *testing.B) {
	s := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(s, ss[i%len(ss)])
		mergesort(s)
	}
}

func BenchmarkBuiltinSort(b *testing.B) {
	s := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(s, ss[i%len(ss)])
		sort.Ints(s)
	}
}

func BenchmarkParallelMergesort1(b *testing.B) {
	s := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(s, ss[i%len(ss)])
		parallelMergesort1(s)
	}
}

func BenchmarkParallelMergesort2(b *testing.B) {
	s := make([]int, size)
	helper := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(s, ss[i%len(ss)])
		parallelMergesort2(s, helper)
	}
}

func BenchmarkParallelMergesort3(b *testing.B) {
	s := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(s, ss[i%len(ss)])
		parallelMergesort3(s)
	}
}
