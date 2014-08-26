library := libblas.a
url := http://www.netlib.org/blas/blas.tgz

build := build
archive := blas.tgz

all: $(library)

$(library): $(build)
	$(MAKE) -C $(build)
	mv $(build)/blas_LINUX.a $@

$(build): $(archive)
	mkdir $@
	tar -xzf $(archive) --strip=1 -C $@

$(archive):
	curl $(url) -o $@

install: $(library)
	go install

clean:
	rm -rf $(archive) $(build) $(library)

.PHONY: all install clean
