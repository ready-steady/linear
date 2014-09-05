library := libblas.a
archive := blas.tgz
url := http://www.netlib.org/blas/$(archive)

build := build

all: $(build)/$(library)

$(build)/$(library): $(build)
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
