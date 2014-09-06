archive := lapack-3.5.0.tgz
url := http://www.netlib.org/lapack/$(archive)

build := build
syso := main.syso

all: $(syso)

install: $(syso)
	go install

$(syso): $(build)/librefblas.a $(build)/liblapack.a
	mkdir -p $(build)/$@
	cd $(build)/$@ && ar x ../librefblas.a
	cd $(build)/$@ && ar x ../liblapack.a
	ld -r -o $@ $(build)/$@/*.o

$(build)/librefblas.a: $(build)
	$(MAKE) -C $(build) blaslib

$(build)/liblapack.a: $(build)
	$(MAKE) -C $(build) lapacklib

$(build): $(archive)
	mkdir -p $@
	tar -xzf $(archive) --strip=1 -C $@

$(archive):
	curl $(url) -o $@

clean:
	rm -rf $(archive) $(build) $(syso)

.PHONY: all install clean
