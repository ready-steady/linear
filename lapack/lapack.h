#ifndef LAPACK_H
#define LAPACK_H

extern void dsyev_(char *JOBZ, char *UPLO, int *N, double *A, int *LDA, double *W,
	double *WORK, int *LWORK, int *INFO);

void dsyev(char JOBZ, char UPLO, int N, double *A, int LDA, double *W,
	double *WORK, int LWORK, int *INFO) {

	dsyev_(&JOBZ, &UPLO, &N, A, &LDA, W, WORK, &LWORK, INFO);
}

#endif
