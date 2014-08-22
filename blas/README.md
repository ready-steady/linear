# BLAS

An adapter for the
[Basic Linear Algebra Subprograms](http://www.netlib.org/blas/). The list of
available routines is limited to those that are utilized by this project, and
it currently includes:

* [dgemv](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html) and
* [dgemm](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html).

## Installation

Run:

```bash
$ go get github.com/gomath/linear/blas
```

The above command might fail. Regardless of the outcome, continue with the
following commands:

```bash
$ cd $GOPATH/src/github.com/gomath/linear/blas
$ make
$ go install
```

The `make` part requires [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries)
to be installed in the system.
