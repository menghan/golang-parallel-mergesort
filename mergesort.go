package mergesort

import (
	"sync"
)

const max = 1 << 11

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

func merge2(s []int, middle int, helper []int) {
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

/* Sequential */

func mergesort(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesort(s[:middle])
		mergesort(s[middle:])
		merge(s, middle)
	}
}

func mergesort2(s []int, helper []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesort2(s[:middle], helper[:middle])
		mergesort2(s[middle:], helper[middle:])
		merge2(s, middle, helper)
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
	len := len(s)

	if len > 1 {
		if len <= max { // Sequential
			mergesort2(s, helper)
		} else { // Parallel
			middle := len / 2

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				parallelMergesort2(s[:middle], helper[:middle])
			}()

			parallelMergesort2(s[middle:], helper[middle:])

			wg.Wait()
			merge2(s, middle, helper)
		}
	}
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
