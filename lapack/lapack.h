#ifndef LAPACK_H
#define LAPACK_H

extern void dgtsv_(const int *N, const int *NRHS, const double *DL,
	double *D, double *DU, double *B, int *LDB, int *INFO);

void dgtsv(int N, int NRHS, double *DL, double *D, double *DU, double *B,
	int LDB, int *INFO) {

	dgtsv_(&N, &NRHS, DL, D, DU, B, &LDB, INFO);
}

extern void dsyev_(const char *JOBZ, const char *UPLO, const int *N,
	double *A, const int *LDA, double *W, double *WORK, const int *LWORK,
	int *INFO);

void dsyev(char JOBZ, char UPLO, int N, double *A, int LDA, double *W,
	double *WORK, int LWORK, int *INFO) {

	dsyev_(&JOBZ, &UPLO, &N, A, &LDA, W, WORK, &LWORK, INFO);
}

#endif
