library := liblapack.a
archive := lapack-3.5.0.tgz
url := http://www.netlib.org/lapack/$(archive)

build := build

all: $(build)/$(library)

$(build)/$(library): $(build)
	$(MAKE) -C $(build) lapacklib

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
