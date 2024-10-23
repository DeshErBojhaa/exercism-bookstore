package bookstore

import "math"

var (
	costs = []int{0, 800, 1600 - 80, 2400 - 240, 3200 - 640, 4000 - 1000}
	dp    = map[struct {
		taken int
		count [5]int
	}]int{}
)

func rec(taken int, cnt [5]int) int {
	if taken == 0 {
		return 0
	}
	key := struct {
		taken int
		count [5]int
	}{taken: taken, count: cnt}
	if ans, ok := dp[key]; ok {
		return ans
	}
	ans := math.MaxInt32
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
	dp[key] = ans
	return ans
}

func Cost(books []int) int {
	cnt := [5]int{}
	for _, book := range books {
		cnt[book-1]++
	}
	return rec(len(books), cnt)
}
