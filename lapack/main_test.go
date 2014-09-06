package lapack

import (
	"math/rand"
	"testing"

	"github.com/go-math/support/assert"
)

func TestDSYEV(t *testing.T) {
	n := 5

	A := []float64{
		0.162182308193243,
		0.601981941401637,
		0.450541598502498,
		0.825816977489547,
		0.106652770180584,
		0.601981941401637,
		0.262971284540144,
		0.083821377996933,
		0.538342435260057,
		0.961898080855054,
		0.450541598502498,
		0.083821377996933,
		0.228976968716819,
		0.996134716626885,
		0.004634224134067,
		0.825816977489547,
		0.538342435260057,
		0.996134716626885,
		0.078175528753184,
		0.774910464711502,
		0.106652770180584,
		0.961898080855054,
		0.004634224134067,
		0.774910464711502,
		0.817303220653433,
	}

	w := make([]float64, n)
	work := []float64{0}
	lwork := -1
	info := 0

	DSYEV('V', 'U', n, A, n, w, work, lwork, &info)

	lwork = int(work[0])
	work = make([]float64, lwork)

	DSYEV('V', 'U', n, A, n, w, work, lwork, &info)

	assert.Equal(info, 0, t)

	eigenVectors := []float64{
		-0.350512137830478,
		+0.116468084895727,
		-0.435005782872646,
		+0.750503447417042,
		-0.333303121372602,
		+0.462361750400701,
		-0.693041256027589,
		-0.409079614137348,
		+0.219801690292016,
		+0.300427221556423,
		-0.638529696902280,
		-0.450088584675982,
		+0.439292414730346,
		+0.201686466663592,
		+0.395025107611391,
		-0.330279135877552,
		+0.306798908225146,
		-0.590447291483772,
		-0.272545278020086,
		+0.611458248553625,
		+0.382829428829298,
		+0.457628341015560,
		+0.319089342511548,
		+0.522946873930437,
		+0.518388356797788,
	}

	eigenValues := []float64{
		-1.145487871954612,
		-0.676875725405419,
		-0.050275996742486,
		+0.892450858666551,
		+2.529798046292787,
	}

	assert.AlmostEqual(A, eigenVectors, t)
	assert.AlmostEqual(w, eigenValues, t)
}

func BenchmarkDSYEV(b *testing.B) {
	n := 1000
	A := genSymmetricMatrix(uint(n))
	w := make([]float64, n)
	lwork := 4 * n
	work := make([]float64, lwork)
	info := 0

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DSYEV('V', 'U', n, A, n, w, work, lwork, &info)
		if info != 0 {
			b.Fatalf("got %v instead of 0", info)
		}
	}
}

func genSymmetricMatrix(n uint) []float64 {
	M := make([]float64, n*n)

	for i := uint(0); i < n; i++ {
		M[i*n+i] = rand.Float64()

		for j := i + 1; j < n; j++ {
			M[i*n+j] = rand.Float64()
			M[j*n+i] = M[i*n+j]
		}
	}

	return M
}
