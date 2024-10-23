// Package bookstore
// Generic solution works for all discounts. Uses dynamic programming.
package bookstore

var (
	costs = []int{0, 800, 1600 - 80, 2400 - 240, 3200 - 640, 4000 - 1000}
	dp    = map[[5]int]int{}
)

func rec(taken int, cnt [5]int) int {
	rem := 0
	for _, c := range cnt {
		rem += c
	}
	if rem == 0 {
		return 0
	}

	if ans, ok := dp[cnt]; ok {
		return ans
	}
	ans := 1 << 31
	for mask := 1; mask < (1 << 5); mask++ {
		n, copyCnt := 0, cnt
		for i := 0; i < 5; i++ {
			if (mask&(1<<i)) == 0 || cnt[i] == 0 {
				continue
			}
			n++
			cnt[i]--
		}
		if n == 0 {
			continue
		}
		tmpAns := rec(taken-n, cnt) + costs[n]
		if tmpAns < ans {
			ans = tmpAns
		}
		cnt = copyCnt
	}
	dp[cnt] = ans
	return ans
}

func Cost(books []int) int {
	cnt := [5]int{}
	for _, book := range books {
		cnt[book-1]++
	}
	return rec(len(books), cnt)
}
