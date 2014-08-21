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

clean:
	rm -rf $(archive) $(build) $(library)

.PHONY: all clean
