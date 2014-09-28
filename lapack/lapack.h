#ifndef LAPACK_H
#define LAPACK_H

extern void dsyev_(char *jobz, char *uplo, int *n, double *a, int *lda, double *w,
	double *work, int *lwork, int *info);

void dsyev(char jobz, char uplo, int n, double *a, int lda, double *w,
	double *work, int lwork, int *info) {

	dsyev_(&jobz, &uplo, &n, a, &lda, w, work, &lwork, info);
}

#endif
