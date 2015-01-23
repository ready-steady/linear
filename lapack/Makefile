build := lapack
syso := main.syso

all: $(syso)

install: $(syso)
	go install

$(syso): $(build)/librefblas.a $(build)/liblapack.a
	mkdir -p $(build)/$@
	cd $(build)/$@ && ar x ../librefblas.a
	cd $(build)/$@ && ar x ../liblapack.a
	ld -r -o $@ $(build)/$@/*.o

$(build)/librefblas.a: $(build)/make.inc
	$(MAKE) -C $(build) blaslib

$(build)/liblapack.a: $(build)/make.inc
	$(MAKE) -C $(build) lapacklib

$(build)/make.inc:
	cp $(build)/make.inc.example $@

clean:
	rm -rf $(build) $(syso)

.PHONY: all install clean
