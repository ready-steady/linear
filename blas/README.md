# BLAS

An interface for the [Basic Linear Algebra
Subprograms](http://www.netlib.org/blas/).

## Installation

If `libblas.a` is available in the system, just run

```bash
$ go get github.com/gomath/linear/blas
```

If `libblas.a` is not available, still run the above command and see it crash.
Then install [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries) and run

```bash
$ cd $GOPATH/pkg/github.com/gomath/linear/blas
$ make
$ go install
```

## Contents

* [dgemv](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html) and
* [dgemm](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html).
