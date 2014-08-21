#ifndef BLAS_H
#define BLAS_H

extern void dgemm_(char *transa, char *transb, int *m, int *n, int *k,
	double *alpha, double *a, int *lda, double *b, int *ldb, double *beta,
	double *c, int *ldc);

#endif
