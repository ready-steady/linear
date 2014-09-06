# LAPACK

The package provides an interface to the
[Linear Algebra PACKage](http://www.netlib.org/lapack/).

The list of available routines currently includes:

* [DSYEV](http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html).

## Installation

Run:

```bash
$ go get github.com/go-math/linear/lapack
```

The above command will fail; however, it will properly clone this repository
into [`$GOPATH`](https://golang.org/doc/code.html#GOPATH). Go to that
directory:

```bash
$ cd $GOPATH/src/github.com/go-math/linal/lapack
```

Then:

```bash
$ make install
```

This command requires [gfortran](https://gcc.gnu.org/wiki/GFortranBinaries) to
be installed in the system.
