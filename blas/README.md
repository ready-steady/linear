# BLAS

An interface for the [Basic Linear Algebra
Subprograms](http://www.netlib.org/blas/). The list of available routines is
restricted to those that are utilized by this project and currently includes:

* [dgemv](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html) and
* [dgemm](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html).

## Installation

If `libblas.a` is available in the system, just run

```bash
$ go get github.com/gomath/linear/blas
```

If `libblas.a` is not available, still run the above command and see it crash.
Then install [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries) and run

```bash
$ cd $GOPATH/src/github.com/gomath/linear/blas
$ make
$ go install
```
