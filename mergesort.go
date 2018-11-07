package mergesort

import (
	"sort"
	"sync"
)

const max = 1 << 11
const mergeMax = 1 << 15

const bubbleMax = 1 << 3

func bubbleSort(s []int) {
	l := len(s)
	if l <= 1 {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < l-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
	}
}

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func search(s []int, v int) int {
	return sort.SearchInts(s, v)
}

func mergeBig(s []int, middle int, helper []int) {
	if len(s) <= 1 || middle == 0 {
		return
	}

	likelyMedianLeft := middle / 2
	likelyMedian := s[likelyMedianLeft]
	likelyMedianRight := middle + search(s[middle:], likelyMedian)

	// trans [0, .A. likelyMedianLeft, .B., middle, .C., likelyMedianRight, .D., len(s)-1]
	// to    [0, .A. likelyMedianLeft, ..C.., ..B.., ..D..]

	newCut := likelyMedianLeft + likelyMedianRight - middle
	if newCut == 0 {
		// A == C == 0, then, [B, D] is in order.
		return
	}

	copy(helper, s[likelyMedianLeft:middle])
	copy(s[likelyMedianLeft:], s[middle:likelyMedianRight])

	var wg sync.WaitGroup
	if likelyMedianLeft != 0 {
		wg.Add(1)
		go func() {
			merge2(s[:newCut], likelyMedianLeft, helper[:newCut])
			wg.Done()
		}()
	}

	copy(s[newCut:], helper[:middle-likelyMedianLeft])
	if middle-likelyMedianLeft != 0 {
		merge2(s[newCut:], middle-likelyMedianLeft, helper[newCut:])
	}

	// log.Printf("From [A=%d B=%d C=%d D=%d] To [{A=%d B=%d} {C=%d D=%d}] s=%+v middle=%d", likelyMedianLeft, middle-likelyMedianLeft, likelyMedianRight-middle, len(s)-likelyMedianRight, likelyMedianLeft, newCut-likelyMedianLeft, likelyMedianRight-newCut, len(s)-likelyMedianRight, s, middle)

	wg.Wait()
}

func merge2(s []int, middle int, helper []int) {
	if len(s) > mergeMax {
		mergeBig(s, middle, helper)
		return
	}

	copy(helper, s[:middle])

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s)

	for helperLeft < middle && helperRight < high {
		if helper[helperLeft] <= s[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = s[helperRight]
			helperRight++
		}
		current++
	}

	if helperLeft < middle {
		copy(s[current:], helper[helperLeft:middle])
	}
}

/* Sequential */

func mergesort(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesort(s[:middle])
		mergesort(s[middle:])
		merge(s, middle)
	}
}

/* Parallel 1 */

func parallelMergesort1(s []int) {
	len := len(s)

	if len > 1 {
		if len <= max { // Sequential
			mergesort(s)
		} else { // Parallel
			middle := len / 2

			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				parallelMergesort1(s[:middle])
			}()

			go func() {
				defer wg.Done()
				parallelMergesort1(s[middle:])
			}()

			wg.Wait()
			merge(s, middle)
		}
	}
}

/* Parallel 2 */

func parallelMergesort2(s []int, helper []int) {
	l := len(s)

	if l <= 1 {
		return
	}

	middle := l / 2

	var wg *sync.WaitGroup

	if l < bubbleMax {
		bubbleSort(s)
		return
	}

	if l > max {
		wg = &sync.WaitGroup{}

		// Parallel
		wg.Add(1)
		go func() {
			parallelMergesort2(s[:middle], helper[:middle])
			wg.Done()
		}()

	} else {
		// Sequential
		parallelMergesort2(s[:middle], helper[:middle])
	}

	parallelMergesort2(s[middle:], helper[middle:])

	if l > max {
		wg.Wait()
	}
	merge2(s, middle, helper)
}

/* Parallel 3 */

func parallelMergesort3(s []int) {
	len := len(s)

	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesort3(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesort3(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}
