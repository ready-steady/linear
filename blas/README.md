# BLAS

The package provides an interface to the
[Basic Linear Algebra Subprograms](http://www.netlib.org/blas/).

The list of available routines currently includes:

* [dgemv](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html) and
* [dgemm](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html).

## Installation

Run:

```bash
$ go get github.com/gomath/linear/blas
```

The above command might fail. Regardless of the outcome, it will properly clone
this repository into [`$GOPATH`](https://golang.org/doc/code.html#GOPATH). Go
to that directory:

```bash
$ cd $GOPATH/src/github.com/gomath/linal/blas
```

Then:

```bash
$ make install
```

This command requires [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries) to
be installed in the system.
