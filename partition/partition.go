package partition

const maxInt = int(^uint(0) >> 1)

func Partition(s []int, k int) [][]int {
	n := len(s)
	d := create2DSlice(n+1, k+1)

	sum := make([]int, n+1)
	sum[0] = 0
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + s[i-1]
	}

	m := create2DSlice(n+1, k+1)
	for i := 1; i <= k; i++ {
		m[1][i] = s[0]
	}
	for i := 2; i <= n; i++ {
		m[i][1] = sum[i]
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			m[i][j] = maxInt

			for x := 1; x <= i-1; x++ {
				cost := max(m[x][j-1], sum[i]-sum[x])

				if cost < m[i][j] {
					m[i][j] = cost
					d[i][j] = x
				}
			}
		}
	}

	return reconstruct(s, d, n, k, [][]int{})
}

func reconstruct(s []int, d [][]int, n int, k int, result [][]int) [][]int {
	if k == 1 {
		result = append(result, s[:n])
	} else {
		result = reconstruct(s, d, d[n][k], k-1, result)
		result = append(result, s[d[n][k]:n])
	}

	return result
}

func create2DSlice(x, y int) [][]int {
	slice := make([][]int, x)

	for i := range slice {
		slice[i] = make([]int, y)
	}

	return slice
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
