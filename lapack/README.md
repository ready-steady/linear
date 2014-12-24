# LAPACK

The package provides an interface to the [Linear Algebra PACKage][lapack], which
also includes the [Basic Linear Algebra Subprograms][blas].

The list of available routines currently includes:

* [DGEMM](http://www.netlib.org/lapack/explore-html/dc/da8/dgemm_8f.html),
* [DGEMV](http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html), and
* [DSYEV](http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html).

## [Documentation][doc]

## Installation

Run:

```bash
$ go get github.com/ready-steady/linear/lapack
```

The above command will fail; however, it will properly clone this repository
into [`$GOPATH`][gopath]. Go to that directory:

```bash
$ cd $GOPATH/src/github.com/ready-steady/linal/lapack
```

Then:

```bash
$ make install
```

This command requires [gfortran][gfortran] to be installed in the system.

[blas]: http://www.netlib.org/blas/
[gfortran]: https://gcc.gnu.org/wiki/GFortranBinaries
[gopath]: https://golang.org/doc/code.html#GOPATH
[lapack]: http://www.netlib.org/lapack/

[doc]: http://godoc.org/github.com/ready-steady/linal/lapack
