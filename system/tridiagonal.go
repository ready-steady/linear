package system

// ComputeTridiagonal solves a tridiagonal system of linear equations.
//
// The system is as follows:
//
//                     b(  0)*x(  0) + c(  1)*x(  1) = d(  0)
//     a(i-1)*x(i-1) + b(  i)*x(  i) + c(i+1)*x(i+1) = d(  i), for i = 1, ..., (n-2),
//     a(n-2)*x(n-2) + b(n-1)*x(n-1)                 = d(n-1)
//
// Note that a(n-1) and c(0) are not involved in the system.
func ComputeTridiagonal(a, b, c, d []float64) []float64 {
	n := len(d)

	cp := make([]float64, n-1)
	cp[0] = c[1] / b[0]
	for i := 1; i < (n - 1); i++ {
		cp[i] = c[i+1] / (b[i] - a[i-1]*cp[i-1])
	}

	dp := make([]float64, n)
	dp[0] = d[0] / b[0]
	for i := 1; i < n; i++ {
		dp[i] = (d[i] - a[i-1]*dp[i-1]) / (b[i] - a[i-1]*cp[i-1])
	}

	x := make([]float64, n)
	x[n-1] = dp[n-1]
	for i := n - 2; i >= 0; i-- {
		x[i] = dp[i] - cp[i]*x[i+1]
	}

	return x
}
