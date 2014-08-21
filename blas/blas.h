#ifndef BLAS_H
#define BLAS_H

extern void dgemv_(char *trans, int *m, int *n, double *alpha, double *a,
	int *lda, double *x, int *incx, double *beta, double *y, int *incy);

extern void dgemm_(char *transa, char *transb, int *m, int *n, int *k,
	double *alpha, double *a, int *lda, double *b, int *ldb, double *beta,
	double *c, int *ldc);

#endif
