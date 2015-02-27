package system

// ComputeTridiagonal solves a tridiagonal system of linear equations.
//
// The system is as follows:
//
//                     b(  0)*x(  0) + c(  1)*x(  1) = d(  0)
//     a(i-1)*x(i-1) + b(  i)*x(  i) + c(i+1)*x(i+1) = d(  i), for i = 1, ..., (n-2),
//     a(n-2)*x(n-2) + b(n-1)*x(n-1)                 = d(n-1)
//
// Note that a(n-1) and c(0) are not involved in the system. The right-hand side
// can be given as a matrix, in which case the system is solved for each column
// of that matrix.
func ComputeTridiagonal(a, b, c, d []float64) []float64 {
	n := len(b)
	nd := len(d) / n

	x := make([]float64, nd*n)

	if n == 1 {
		for j := 0; j < nd; j++ {
			x[j] = d[j] / b[0]
		}
		return x
	}

	bp := make([]float64, n-1)
	cp := make([]float64, n-1)
	cp[0] = c[1] / b[0]
	bp[0] = b[1] - a[0]*cp[0]
	for i := 1; i < (n - 1); i++ {
		cp[i] = c[i+1] / bp[i-1]
		bp[i] = b[i+1] - a[i]*cp[i]
	}

	dp := make([]float64, n)
	for i := 0; i < nd; i++ {
		dp[0] = d[i*n] / b[0]
		for j := 1; j < n; j++ {
			dp[j] = (d[i*n+j] - a[j-1]*dp[j-1]) / bp[j-1]
		}

		x[i*n+n-1] = dp[n-1]
		for j := n - 2; j >= 0; j-- {
			x[i*n+j] = dp[j] - cp[j]*x[i*n+j+1]
		}
	}

	return x
}
