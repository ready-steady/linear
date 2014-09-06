library := liblapack.a
archive := lapack-3.5.0.tgz
url := http://www.netlib.org/lapack/$(archive)

build := build

all: $(library)

$(library): $(build)
	$(MAKE) -C $(build) blaslib lapacklib
	cp $(build)/librefblas.a .
	cp $(build)/liblapack.a $@

$(build): $(archive)
	mkdir $@
	tar -xzf $(archive) --strip=1 -C $@

$(archive):
	curl $(url) -o $@

install: $(library)
	go install

clean:
	rm -rf $(archive) $(build) *.a

.PHONY: all install clean
