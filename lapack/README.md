# LAPACK

The package provides an interface to the [Linear Algebra
PACKage](http://www.netlib.org/lapack/), which also includes the [Basic Linear
Algebra Subprograms](http://www.netlib.org/blas/).

The list of available routines currently includes:

* [DGEMM](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html),
* [DGEMV](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html), and
* [DSYEV](http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html).

## Installation

Run:

```bash
$ go get github.com/ready-steady/linear/lapack
```

The above command will fail; however, it will properly clone this repository
into [`$GOPATH`](https://golang.org/doc/code.html#GOPATH). Go to that
directory:

```bash
$ cd $GOPATH/src/github.com/ready-steady/linal/lapack
```

Then:

```bash
$ make install
```

This command requires [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries) to
be installed in the system.
