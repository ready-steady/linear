#ifndef BLAS_H
#define BLAS_H

extern void dgemv_(const char *TRANS, const int *M, const int *N,
	const double *ALPHA, const double *A, const int *LDA, const double *X,
	const int *INCX, const double *BETA, double *Y, const int *INCY);

void dgemv(char TRANS, int M, int N, double ALPHA, const double *A,
	int LDA, const double *X, int INCX, double BETA, double *Y, int INCY) {

	dgemv_(&TRANS, &M, &N, &ALPHA, A, &LDA, X, &INCX, &BETA, Y, &INCY);
}

extern void dgemm_(const char *TRANSA, const char *TRANSB, const int *M,
	const int *N, const int *K, const double *ALPHA, const double *A,
	const int *LDA, const double *B, const int *LDB, const double *BETA,
	double *C, const int *LDC);

void dgemm(char TRANSA, char TRANSB, int M, int N, int K, double ALPHA,
	const double *A, int LDA, const double *B, int LDB, double BETA,
	double *C, int LDC) {

	dgemm_(&TRANSA, &TRANSB, &M, &N, &K, &ALPHA, A, &LDA, B, &LDB, &BETA, C, &LDC);
}

#endif
