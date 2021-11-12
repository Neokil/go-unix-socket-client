install:
	go build -o gusc
	chmod a+x gusc
	ln -sf $(shell pwd)/gusc /usr/local/bin/gusc
